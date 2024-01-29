package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/microcosm-cc/bluemonday"
	"net/http"
	"open-rev.com/helper"
	"open-rev.com/infrastructure/dto"
	"open-rev.com/usecase"
	"strings"
	"time"
)

type areaHandler struct {
	Contract    client.Contract
	AreaUsecase usecase.AreaUsecase
}

type AreaHandler interface {
	GetAllAreas(ctx *gin.Context)
	GetAllSubAreas(ctx *gin.Context)
	GetAllAreasAndSubAreas(ctx *gin.Context)
	DeleteArea(ctx *gin.Context)
	DeleteSubArea(ctx *gin.Context)
	AddArea(ctx *gin.Context)
	AddSubArea(ctx *gin.Context)
}

func NewAreaHandler(areaUsecase usecase.AreaUsecase, contract *client.Contract) AreaHandler {
	return &areaHandler{AreaUsecase: areaUsecase, Contract: *contract}
}

func (a *areaHandler) AddArea(ctx *gin.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)

	var newArea dto.AddAreaDto
	if err := decoder.Decode(&newArea); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": helper.Body_decoding_err})
		return
	}

	policy := bluemonday.UGCPolicy()
	newArea.Name = strings.TrimSpace(policy.Sanitize(newArea.Name))

	if newArea.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": helper.XSS})
		return
	}
	newArea.LastUpdateTime = time.Now()
	err := a.AreaUsecase.AddArea(ctx, a.Contract, newArea)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func (a *areaHandler) AddSubArea(ctx *gin.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)

	var newSubarea dto.AddSubAreaDto
	if err := decoder.Decode(&newSubarea); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": helper.Body_decoding_err})
		return
	}

	policy := bluemonday.UGCPolicy()
	newSubarea.Name = strings.TrimSpace(policy.Sanitize(newSubarea.Name))
	newSubarea.AreaId = strings.TrimSpace(policy.Sanitize(newSubarea.AreaId))
	newSubarea.LastUpdateTime = time.Now()
	if newSubarea.Name == "" || newSubarea.AreaId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": helper.XSS})
		return
	}

	err := a.AreaUsecase.AddSubArea(ctx, a.Contract, newSubarea)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
}

// todo: handle deleted lastupdatetim
func (a *areaHandler) DeleteArea(ctx *gin.Context) {
	id := ctx.Param("id")

	err := a.AreaUsecase.DeleteArea(ctx, a.Contract, id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully deleted"})
}

// todo: handle deleted lastupdatetim

func (a *areaHandler) DeleteSubArea(ctx *gin.Context) {
	id := ctx.Param("id")

	err := a.AreaUsecase.DeleteSubArea(ctx, a.Contract, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully deleted"})
}

func (a *areaHandler) GetAllAreasAndSubAreas(ctx *gin.Context) {
	areas, err := a.AreaUsecase.GetAllAreasAndSubAreas(ctx, a.Contract)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, areas)
}

func (a *areaHandler) GetAllAreas(ctx *gin.Context) {
	areas, err := a.AreaUsecase.GetAllAreas(ctx, a.Contract)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, areas)
}

func (a *areaHandler) GetAllSubAreas(ctx *gin.Context) {
	subareas, err := a.AreaUsecase.GetAllSubAreas(ctx, a.Contract)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, subareas)
}
