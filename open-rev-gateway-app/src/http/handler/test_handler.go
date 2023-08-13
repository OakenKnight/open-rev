package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"net/http"
	"open-rev.com/usecase"
)

type testHandler struct {
	Contract    client.Contract
	TestUsecase usecase.TestUsecase
}

func (t testHandler) TestCCVersion(ctx *gin.Context) {
	version, err := t.TestUsecase.TestVersion(ctx, t.Contract)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, version)
}

type TestHandler interface {
	TestCCVersion(ctx *gin.Context)
}

func NewTestHandler(testUsecase usecase.TestUsecase, contract *client.Contract) TestHandler {
	return &testHandler{TestUsecase: testUsecase, Contract: *contract}
}
