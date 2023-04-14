package main

import (
	"context"
	"crypto/x509"
	"fmt"
	"github.com/minio/minio-go/v7"
	"io/ioutil"
	"log"
	"open-rev.com/config"
	"open-rev.com/http/router"
	"open-rev.com/interactor"
	"path"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	minioCredentials "github.com/minio/minio-go/v7/pkg/credentials"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	log.Println("============ application-golang starts ============")

	cfg, err := config.GetConfig("./.env")
	if err != nil {
		panic(err)
	}

	// The gRPC client connection should be shared by all Gateway connections to this endpoint
	clientConnection := newGrpcConnection(cfg)
	defer clientConnection.Close()

	id := newIdentity(cfg)
	sign := newSign(cfg)

	// Create a Gateway connection for a specific client identity
	gateway, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		// Default timeouts for different gRPC calls
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		panic(err)
	}
	defer gateway.Close()

	network := gateway.GetNetwork(cfg.ChannelName)
	contract := network.GetContract(cfg.ChaincodeName)

	//if cfg.Mode == "development" && cfg.IsCompose == true {
	log.Println("Initializing ledger:")
	initLedger(contract)
	//}

	// TODO: Implement fail safe mechanism with exponential retries in case of pod being down
	minioClient, err := minio.New(cfg.MinioUrl, &minio.Options{
		Creds:  minioCredentials.NewStaticV4(cfg.MinioAccessId, cfg.MinioSecret, ""),
		Secure: cfg.MinioSsl,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", minioClient)
	bucketName := cfg.MinioBucket
	location := cfg.MinioLocation
	ctx := context.Background()

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	i := interactor.NewInteractor(*contract, *minioClient)
	appHandler := i.NewAppHandler()

	gin.SetMode(cfg.GinMode)
	r := router.NewRouter(appHandler, contract)
	r.Use(gin.Logger())

	// TODO: further fix on CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	// establish minio connection
	// TODO: Implement fail safe mechanism with exponential retries in case of pod being down

	//// Upload the zip file
	//objectName := "PDAJ-Projekat.pdf"
	//filePath := "../PDAJ-Projekat.pdf"
	//contentType := "application/pdf"

	// Upload the zip file with FPutObject
	//info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
	//
	//object, err := minioClient.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer object.Close()
	//
	//localFile, err := os.Create("../local-file.pdf")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer localFile.Close()
	//
	//if _, err = io.Copy(localFile, object); err != nil {
	//	fmt.Println(err)
	//	return
	//}

	// TODO: run on TLS depending on env variable
	if cfg.Mode == "Production" {
		err = r.RunTLS(":443", cfg.GinCertPath, cfg.GinKeyPath)
	} else {
		err = r.Run(":" + cfg.Port)
	}

	if err != nil {
		return
	}

}

// newIdentity creates a client identity for this Gateway connection using an X.509 certificate.
func newIdentity(config *config.Config) *identity.X509Identity {

	certificate, err := loadCertificate(config.CertPath)
	if err != nil {
		panic(err)
	}

	id, err := identity.NewX509Identity(config.MspID, certificate)
	if err != nil {
		panic(err)
	}

	return id
}

// newGrpcConnection creates a gRPC connection to the Gateway server.
func newGrpcConnection(config *config.Config) *grpc.ClientConn {
	certificate, err := loadCertificate(config.TlsCertPath)
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, config.GatewayPeer)

	connection, err := grpc.Dial(config.PeerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		panic(fmt.Errorf("failed to create gRPC connection: %w", err))
	}

	return connection
}

func loadCertificate(filename string) (*x509.Certificate, error) {
	certificatePEM, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file: %w", err)
	}
	return identity.CertificateFromPEM(certificatePEM)
}

// This type of transaction would typically only be run once by an application the first time it was started after its
// initial deployment. A new version of the chaincode deployed later would likely not need to run an "init" function.
func initLedger(contract *client.Contract) {
	log.Printf("Submit Transaction: InitLedger, function creates the initial set of assets on the ledger \n")

	_, err := contract.SubmitTransaction("InitLedger")
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction: %w", err))
	}

	log.Printf("*** Transaction committed successfully ***\n")
}

// newSign creates a function that generates a digital signature from a message digest using a private key.
func newSign(config *config.Config) identity.Sign {
	files, err := ioutil.ReadDir(config.KeyPath)
	if err != nil {
		panic(fmt.Errorf("failed to read private key directory: %w", err))
	}
	privateKeyPEM, err := ioutil.ReadFile(path.Join(config.KeyPath, files[0].Name()))

	if err != nil {
		panic(fmt.Errorf("failed to read private key file: %w", err))
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		panic(err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		panic(err)
	}

	return sign
}
