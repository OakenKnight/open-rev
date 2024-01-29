package handler

import (
	"encoding/json"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/microcosm-cc/bluemonday"
	"open-rev.com/domain"
	"open-rev.com/helper"
	"open-rev.com/infrastructure/dto"
	"open-rev.com/usecase"
)

type userHandler struct {
	Contract    client.Contract
	UserUsecase usecase.UserUsecase
}

func (u *userHandler) GetTopReviewers(ctx *gin.Context) {
	users, err := u.UserUsecase.GetAllUsersWithDetails(ctx, u.Contract)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	sort.Slice(users, func(i, j int) bool {
		return users[i].AvgMyRevsQuality > users[j].AvgMyRevsQuality
	})
	var top []*dto.TopReviewerDto
	for _, user := range users {
		top = append(top, &dto.TopReviewerDto{AvgReview: user.AvgMyRevsQuality, User: user.Name + " " + user.Surname, Guid: user.ID, Email: user.Email, RoleId: user.RoleId})
	}
	ctx.JSON(http.StatusOK, top)
}
func (u *userHandler) GetTopAuthors(ctx *gin.Context) {
	users, err := u.UserUsecase.GetAllUsersWithDetails(ctx, u.Contract)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	sort.Slice(users, func(i, j int) bool {
		return users[i].AvgMark > users[j].AvgMark
	})
	var top []*dto.TopAuthorDto
	for _, user := range users {
		top = append(top, &dto.TopAuthorDto{AvgRate: user.AvgMark, User: user.Name + " " + user.Surname, Guid: user.ID, Email: user.Email, RoleId: user.RoleId})
	}
	ctx.JSON(http.StatusOK, top)
}

// todo: handle deleted lastupdatetim
func (u *userHandler) DeleteOpenRevUser(ctx *gin.Context) {
	id := ctx.Param("id")

	err := u.UserUsecase.DeleteOpenRevUser(ctx, u.Contract, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully deleted"})
}

func (u *userHandler) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := u.UserUsecase.GetUserById(ctx, u.Contract, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (u *userHandler) EditUser(ctx *gin.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var user dto.EditUserDTO
	if err := decoder.Decode(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": helper.Body_decoding_err})
		return
	}

	policy := bluemonday.UGCPolicy()
	user.ID = strings.TrimSpace(policy.Sanitize(user.ID))
	user.Name = strings.TrimSpace(policy.Sanitize(user.Name))
	user.Surname = strings.TrimSpace(policy.Sanitize(user.Surname))

	if user.Name == "" || user.Surname == "" || user.ID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": helper.XSS})
		return
	}
	user.LastUpdateTime = time.Now()
	err := u.UserUsecase.EditUser(ctx, u.Contract, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func (u *userHandler) ConfirmAccount(ctx *gin.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var confirmDto dto.ConfirmAccountDTO
	if err := decoder.Decode(&confirmDto); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": helper.Body_decoding_err})
		return
	}

	policy := bluemonday.UGCPolicy()
	confirmDto.ID = strings.TrimSpace(policy.Sanitize(confirmDto.ID))
	confirmDto.Code = strings.TrimSpace(policy.Sanitize(confirmDto.Code))

	if confirmDto.ID == "" || confirmDto.Code == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": helper.XSS})
		return
	}

	newRegisteredUser, err := u.UserUsecase.ConfirmAccount(ctx, u.Contract, confirmDto.Code, confirmDto.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, *newRegisteredUser)
}

func (u *userHandler) Register(ctx *gin.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var user domain.OpenRevUser
	if err := decoder.Decode(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": helper.Body_decoding_err})
		return
	}

	policy := bluemonday.UGCPolicy()
	user.ID = strings.TrimSpace(policy.Sanitize(user.ID))
	user.Name = strings.TrimSpace(policy.Sanitize(user.Name))
	user.Surname = strings.TrimSpace(policy.Sanitize(user.Surname))
	user.Email = strings.TrimSpace(policy.Sanitize(user.Email))

	if user.Name == "" || user.Surname == "" || user.Email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": helper.XSS})
		return
	}
	newRegisteredUser, err := u.UserUsecase.Register(ctx, u.Contract, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, *newRegisteredUser)
}
func (u *userHandler) GetUserInfo(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := u.UserUsecase.GetUserInfo(ctx, u.Contract, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
func (u *userHandler) GetAllUsers(ctx *gin.Context) {
	users, err := u.UserUsecase.GetAllUsers(ctx, u.Contract)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

type UserHandler interface {
	EditUser(ctx *gin.Context)
	GetAllUsers(ctx *gin.Context)
	Register(ctx *gin.Context)
	ConfirmAccount(ctx *gin.Context)
	GetUserById(ctx *gin.Context)
	GetUserInfo(ctx *gin.Context)
	DeleteOpenRevUser(ctx *gin.Context)
	GetTopReviewers(ctx *gin.Context)
	GetTopAuthors(ctx *gin.Context)
}

func NewUserHandler(userUsecase usecase.UserUsecase, contract *client.Contract) UserHandler {
	return &userHandler{UserUsecase: userUsecase, Contract: *contract}
}
