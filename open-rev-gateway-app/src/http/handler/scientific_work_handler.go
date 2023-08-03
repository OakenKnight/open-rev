package handler

import (
	"fmt"
	"github.com/minio/minio-go/v7"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"open-rev.com/config"
	"os"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/microcosm-cc/bluemonday"
	"open-rev.com/helper"
	"open-rev.com/usecase"
)

type scientificWorkHandler struct {
	Contract              client.Contract
	MinioClient           minio.Client
	ScientificWorkUsecase usecase.ScientificWorkUsecase
}

func (s *scientificWorkHandler) DownloadScientificWork(ctx *gin.Context) {
	filename := ctx.Param("filename")
	bucket := ctx.Param("bucket")
	minioObj, err := s.MinioClient.GetObject(ctx, bucket, filename, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "File reading from MinIO error"})
		return
	}
	defer minioObj.Close()
	b := make([]byte, 5000000)
	num, err := minioObj.Read(b)
	permissions := 0644 // or whatever you need

	err1 := ioutil.WriteFile("file.pdf", b, fs.FileMode(permissions))
	if err1 != nil {
		// handle error
	}

	if err.Error() == "EOF" {
		ctx.Data(http.StatusOK, "application/octet-stream", b[0:num])
		return
	}
}

const (
	contentType = "application/pdf"
)

type ScientificWorkHandler interface {
	DeleteScientificWork(ctx *gin.Context)
	GetAllScientificWorks(ctx *gin.Context)
	CreateScientificWork(ctx *gin.Context)
	GetAllScientificWorksByUser(context *gin.Context)
	GetScientificWorkById(context *gin.Context)
	GetDashboard(context *gin.Context)
	GetAllScientificWorksBySubareaId(context *gin.Context)
	GetScientificWorkDetails(context *gin.Context)
	GetTopScientificWorks(ctx *gin.Context)
	DownloadScientificWork(ctx *gin.Context)
}

func NewScientificWorkHandler(scientificWorkUsecase usecase.ScientificWorkUsecase, contract *client.Contract, minioClient *minio.Client) ScientificWorkHandler {
	return &scientificWorkHandler{ScientificWorkUsecase: scientificWorkUsecase, Contract: *contract, MinioClient: *minioClient}
}

func (s *scientificWorkHandler) DeleteScientificWork(ctx *gin.Context) {
	id := ctx.Param("id")

	err := s.ScientificWorkUsecase.DeleteSciWork(ctx, s.Contract, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully deleted"})
}

func (s *scientificWorkHandler) GetAllScientificWorksByUser(ctx *gin.Context) {
	userId := ctx.Param("userId")
	works, err := s.ScientificWorkUsecase.GetAllScientificWorksByUser(ctx, s.Contract, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, works)
}

func (s *scientificWorkHandler) GetAllScientificWorksBySubareaId(ctx *gin.Context) {
	subareaId := ctx.Param("subareaId")
	works, err := s.ScientificWorkUsecase.GetAllScientificWorksBySubareaId(ctx, s.Contract, subareaId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if works == nil {
		ctx.JSON(http.StatusNotFound, make([]string, 0))
		return
	}
	ctx.JSON(http.StatusOK, works)
}

func (s *scientificWorkHandler) GetScientificWorkById(ctx *gin.Context) {
	id := ctx.Param("id")
	works, err := s.ScientificWorkUsecase.GetScientificWorkById(ctx, s.Contract, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, works)

}
func (s *scientificWorkHandler) GetScientificWorkDetails(ctx *gin.Context) {
	id := ctx.Param("id")
	works, err := s.ScientificWorkUsecase.GetScientificWorkDetails(ctx, s.Contract, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, works)

}

func (s *scientificWorkHandler) GetDashboard(ctx *gin.Context) {
	dash, err := s.ScientificWorkUsecase.GetDashboard(ctx, s.Contract)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dash)

}

func (s *scientificWorkHandler) GetAllScientificWorks(ctx *gin.Context) {
	works, err := s.ScientificWorkUsecase.GetAllScientificWorks(ctx, s.Contract)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, works)
}

func (s *scientificWorkHandler) GetTopScientificWorks(ctx *gin.Context) {
	sciworks, err := s.ScientificWorkUsecase.GetAllScientificWorksWithDetails(ctx, s.Contract)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	sort.Slice(sciworks, func(i, j int) bool {
		return sciworks[i].AvgRate > sciworks[j].AvgRate
	})

	ctx.JSON(http.StatusOK, sciworks)
}
func (s *scientificWorkHandler) CreateScientificWork(ctx *gin.Context) {
	sciWork := helper.TransformForm(ctx)
	policy := bluemonday.UGCPolicy()

	sciWork.Title = strings.TrimSpace(policy.Sanitize(sciWork.Title))
	sciWork.SubAreaId = strings.TrimSpace(policy.Sanitize(sciWork.SubAreaId))
	sciWork.Abstract = strings.TrimSpace(policy.Sanitize(sciWork.Abstract))
	sciWork.Keywords = strings.TrimSpace(policy.Sanitize(sciWork.Keywords))
	sciWork.UserId = strings.TrimSpace(policy.Sanitize(sciWork.UserId))

	if sciWork.Title == "" || sciWork.SubAreaId == "" || sciWork.Abstract == "" || sciWork.Keywords == "" || sciWork.UserId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": helper.XSS})
		return
	}

	file, h, err := ctx.Request.FormFile("PdfFile")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error reading file"})
		return
	}

	fileUuid, err := uuid.NewUUID()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error creating uuid for saved file "})
		return
	}

	cfg, err := config.GetConfig("./.env")
	if err != nil {
		panic(err)
	}
	filePath := ""
	h.Filename = strings.ReplaceAll(h.Filename, " ", "")
	var out *os.File
	if cfg.IsCompose {
		out, err = os.Create("src/public/" + h.Filename)
		filePath = "src/public/"
	} else {
		out, err = os.Create("public/" + h.Filename)
		filePath = "public/"
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error opening new file"})
		return
	}

	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "File saving error"})
		return
	}
	filename := fileUuid.String() + "_" + h.Filename

	info, err := s.MinioClient.FPutObject(ctx, cfg.MinioBucket, filename, filePath+h.Filename, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "File saving to MinIO error"})
		return
	}
	log.Printf("[MinIO][OK]:Successfully uploaded %s of size %d\n", filename, info.Size)

	sciWork.PdfFile = cfg.MinioBucket + "/" + filename

	wd, err := os.Getwd()
	var fileUri string
	if cfg.IsCompose {
		fileUri = wd + "/src/public/" + filename
	} else {
		fileUri = wd + "/public/" + filename
	}

	err = os.Remove(fileUri)
	if err != nil {
		log.Println(err)
	}

	newWork, err := s.ScientificWorkUsecase.CreateScientificWork(ctx, s.Contract, &sciWork)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, newWork)
	return
}
