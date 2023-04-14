package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"open-rev.com/infrastructure/dto"
)

func FormatJSON(data []byte) string {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, data, " ", ""); err != nil {
		panic(fmt.Errorf("failed to parse JSON: %w", err))
	}
	return prettyJSON.String()
}

func TransformForm(ctx *gin.Context) dto.NewScientificWorkDTO {
	var newSciWork dto.NewScientificWorkDTO
	newSciWork.Title = ctx.PostForm("Title")
	newSciWork.Keywords = ctx.PostForm("Keywords")
	newSciWork.UserId = ctx.PostForm("UserId")
	newSciWork.SubAreaId = ctx.PostForm("SubAreaId")
	newSciWork.Abstract = ctx.PostForm("Abstract")
	return newSciWork
}
