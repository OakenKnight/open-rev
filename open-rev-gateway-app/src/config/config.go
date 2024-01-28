package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Mode            string
	MspID           string
	CryptoPath      string
	CertPath        string
	KeyPath         string
	TlsCertPath     string
	PeerEndpoint    string
	GatewayPeer     string
	ChannelName     string
	ChaincodeName   string
	GinMode         string
	GinCertPath     string
	GinKeyPath      string
	Port            string
	SeaweedFsAssign string
	IsCompose       bool
	MinioAccessId   string
	MinioSecret     string
	MinioUrl        string
	MinioSsl        bool
	MinioBucket     string
	MinioLocation   string
	DefaultPath     string
}

func GetConfig(path string) (*Config, error) {
	config := Config{}
	//Uncomment for local development
	if os.Getenv("IS_COMPOSE") != "true" {
		err := godotenv.Load(path)
		if err != nil {
			return nil, fmt.Errorf("error loading .env file")
		}
	}

	config.Mode = os.Getenv("MODE")
	config.MspID = os.Getenv("MSP_ID")
	config.CryptoPath = os.Getenv("CRYPTO_PATH")
	config.KeyPath = config.CryptoPath + os.Getenv("KEY_PATH")
	config.CertPath = config.CryptoPath + os.Getenv("CERT_PATH")
	config.TlsCertPath = config.CryptoPath + os.Getenv("TLS_CERT_PATH")

	config.PeerEndpoint = os.Getenv("PEER_ENDPOINT_DEVELOP")
	config.GatewayPeer = os.Getenv("GATEWAY_PEER_DEVELOP")
	config.ChannelName = os.Getenv("CHANNEL_NAME")
	config.ChaincodeName = os.Getenv("CHAINCODE_NAME")
	config.GinMode = os.Getenv("GIN_MODE")
	config.GinCertPath = os.Getenv("CERT")
	config.GinKeyPath = os.Getenv("KEY")
	config.Port = os.Getenv("PORT")
	config.SeaweedFsAssign = os.Getenv("SEAWEED_FS_PATH_ASSIGN")
	config.MinioAccessId = os.Getenv("MINIO_ACCESS_ID")
	config.MinioSecret = os.Getenv("MINIO_SECRET")
	config.MinioUrl = os.Getenv("MINIO_URL")
	config.MinioBucket = os.Getenv("MINIO_BUCKET")
	config.MinioLocation = os.Getenv("MINIO_LOCATION")
	config.MinioSsl = os.Getenv("Minio_SSL") == "true"
	config.DefaultPath = os.Getenv("DEFAULT_PATH")
	config.IsCompose = os.Getenv("IS_COMPOSE") == "true"
	return &config, nil
}
