package interactor

import (
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/minio/minio-go/v7"
	"io"
	"open-rev.com/http/handler"
	"open-rev.com/usecase"
)

type interactor struct {
	Closer      io.Closer
	Contract    client.Contract
	MinioClient minio.Client
}

type Interactor interface {
	NewAppHandler() AppHandler
}

type appHandler struct {
	handler.AreaHandler
	handler.UserHandler
	handler.ReviewHandler
	handler.ScientificWorkHandler
	handler.TestHandler
}

type AppHandler interface {
	handler.AreaHandler
	handler.UserHandler
	handler.ReviewHandler
	handler.ScientificWorkHandler
	handler.TestHandler
}

func NewInteractor(contract client.Contract, minioClient minio.Client) Interactor {
	return &interactor{Contract: contract, MinioClient: minioClient}
}

func (i *interactor) NewAppHandler() AppHandler {
	appHandler := &appHandler{}
	appHandler.AreaHandler = i.NewAreaHandler(i.Contract)
	appHandler.UserHandler = i.NewUserHandler(i.Contract)
	appHandler.ScientificWorkHandler = i.NewScientificWorkHandler(i.Contract, i.MinioClient)
	appHandler.ReviewHandler = i.NewReviewHandler(i.Contract)
	appHandler.TestHandler = i.NewTestHandler(i.Contract)
	return appHandler
}

func (i *interactor) NewAreaUsecase(contract client.Contract) usecase.AreaUsecase {
	return usecase.NewAreaUsecase(&contract)
}

func (i *interactor) NewAreaHandler(contract client.Contract) handler.AreaHandler {
	return handler.NewAreaHandler(i.NewAreaUsecase(contract), &contract)
}

func (i *interactor) NewReviewUsecase(contract client.Contract) usecase.ReviewUsecase {
	return usecase.NewReviewUsecase(&contract)
}

func (i *interactor) NewReviewHandler(contract client.Contract) handler.ReviewHandler {
	return handler.NewReviewHandler(i.NewReviewUsecase(contract), &contract)
}

func (i *interactor) NewUserUsecase(contract client.Contract) usecase.UserUsecase {
	return usecase.NewUserUsecase(&contract)
}

func (i *interactor) NewUserHandler(contract client.Contract) handler.UserHandler {
	return handler.NewUserHandler(i.NewUserUsecase(contract), &contract)
}

func (i *interactor) NewScientificWorkHandler(contract client.Contract, minioClient minio.Client) handler.ScientificWorkHandler {
	return handler.NewScientificWorkHandler(i.NewScientificWorkUsecase(contract), &contract, &minioClient)

}

func (i *interactor) NewScientificWorkUsecase(contract client.Contract) usecase.ScientificWorkUsecase {
	return usecase.NewScientificWorkUsecase(&contract)
}
func (i *interactor) NewTestUsecase(contract client.Contract) usecase.TestUsecase {
	return usecase.NewTestUsecase(&contract)
}

func (i *interactor) NewTestHandler(contract client.Contract) handler.TestHandler {
	return handler.NewTestHandler(i.NewTestUsecase(contract), &contract)
}
