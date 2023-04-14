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
)

type reviewHandler struct {
	Contract      client.Contract
	ReviewUsecase usecase.ReviewUsecase
}

type ReviewHandler interface {
	GetAllReviews(ctx *gin.Context)
	GetAllReviewsByScientificWork(ctx *gin.Context)
	GetAllReviewsByUser(ctx *gin.Context)
	CreateReview(ctx *gin.Context)
	CreateReviewQuality(ctx *gin.Context)
	DeleteReview(ctx *gin.Context)
}

func NewReviewHandler(reviewUsecase usecase.ReviewUsecase, contract *client.Contract) ReviewHandler {
	return &reviewHandler{ReviewUsecase: reviewUsecase, Contract: *contract}
}
func (r *reviewHandler) DeleteReview(ctx *gin.Context) {
	id := ctx.Param("id")

	err := r.ReviewUsecase.DeleteReview(ctx, r.Contract, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully deleted"})
}

func (u *reviewHandler) CreateReviewQuality(ctx *gin.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var review dto.ReviewQualityDTO
	if err := decoder.Decode(&review); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": helper.Body_decoding_err})
		return
	}

	policy := bluemonday.UGCPolicy()

	review.ReviewId = strings.TrimSpace(policy.Sanitize(review.ReviewId))
	review.UserId = strings.TrimSpace(policy.Sanitize(review.UserId))

	if review.ReviewId == "" || review.UserId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": helper.XSS})
		return
	}

	_, err := u.ReviewUsecase.CreateReviewQuality(ctx, u.Contract, &review)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success_msg": "Successfully inserted review quality mark!"})
}

func (u *reviewHandler) GetAllReviews(ctx *gin.Context) {
	reviews, err := u.ReviewUsecase.GetAllReviews(ctx, u.Contract)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, reviews)
}

func (u *reviewHandler) CreateReview(ctx *gin.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var review dto.ReviewDTO
	if err := decoder.Decode(&review); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": helper.Body_decoding_err})
		return
	}

	policy := bluemonday.UGCPolicy()

	review.Review = strings.TrimSpace(policy.Sanitize(review.Review))
	review.ScientificWorkId = strings.TrimSpace(policy.Sanitize(review.ScientificWorkId))
	review.UserId = strings.TrimSpace(policy.Sanitize(review.UserId))

	if review.Review == "" || review.ScientificWorkId == "" || review.UserId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": helper.XSS})
		return
	}

	_, err := u.ReviewUsecase.CreateReview(ctx, u.Contract, review)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success_msg": "Successfully inserted review!"})
}

func (u *reviewHandler) GetAllReviewsByScientificWork(ctx *gin.Context) {
	id := ctx.Param("id")
	reviews, err := u.ReviewUsecase.GetAllReviewsByScientificWork(ctx, u.Contract, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, reviews)
}

func (u *reviewHandler) GetAllReviewsByUser(ctx *gin.Context) {
	id := ctx.Param("id")
	reviews, err := u.ReviewUsecase.GetAllReviewsByUser(ctx, u.Contract, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, reviews)
}
