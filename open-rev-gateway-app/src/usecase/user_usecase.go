package usecase

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"log"
	"open-rev.com/domain"
	"open-rev.com/helper"
	"open-rev.com/infrastructure/dto"
)

type userUsecase struct {
	Contract *client.Contract
}

func (r *userUsecase) GetAllUsersWithDetails(context context.Context, contract client.Contract) ([]*dto.OpenRevUserInfoDTO, error) {
	log.Printf("Evaluate Transaction: ReadAllUsersWithDetails, function returns users with details on the ledger")

	evaluateResult, err := contract.EvaluateTransaction("ReadAllUsersWithDetails")
	if err != nil {
		return nil, helper.LedgerErrorHandler(&contract, err)
	}

	var users []*dto.OpenRevUserInfoDTO

	err = json.Unmarshal(evaluateResult, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userUsecase) DeleteOpenRevUser(context context.Context, contract client.Contract, id string) error {
	log.Println("Submit Transaction: DeleteOpenRevUserAsset, function returns error if not successful")
	_, err := contract.SubmitTransaction("DeleteOpenRevUserAsset", id)
	if err != nil {
		return helper.LedgerErrorHandler(&contract, err)
	}
	return nil
}

func (r *userUsecase) GetUserInfo(context context.Context, contract client.Contract, id string) (*dto.OpenRevUserInfoDTO, error) {
	log.Printf("Evaluate Transaction: ReadOpenRevUserInfoAsset, function returns user %s on the ledger", id)

	evaluateResult, err := contract.EvaluateTransaction("ReadOpenRevUserInfoAsset", id)
	if err != nil {
		return nil, helper.LedgerErrorHandler(&contract, err)
	}

	var user *dto.OpenRevUserInfoDTO

	err = json.Unmarshal(evaluateResult, &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userUsecase) GetUserById(context context.Context, contract client.Contract, id string) (*domain.OpenRevUser, error) {
	log.Printf("Evaluate Transaction: ReadOpenRevUserAsset, function returns users asset %s on the ledger", id)

	evaluateResult, err := contract.EvaluateTransaction("ReadOpenRevUserAsset", id)
	if err != nil {
		return nil, helper.LedgerErrorHandler(&contract, err)
	}

	var user *domain.OpenRevUser

	err = json.Unmarshal(evaluateResult, &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userUsecase) EditUser(context context.Context, contract client.Contract, user dto.EditUserDTO) error {
	log.Printf("Submit Transaction: EditOpenRevUserAsset, function edits user with id %s", user.ID)

	evaluateResult, err := contract.SubmitTransaction("EditOpenRevUserAsset", user.ID, user.Name, user.Surname)
	if err != nil {
		return helper.LedgerErrorHandler(&contract, err)
	}

	var openRevUser domain.OpenRevUser
	err = json.Unmarshal(evaluateResult, &openRevUser)
	if err != nil {
		return err
	}

	return nil
}

func (r *userUsecase) GetAllUsers(context context.Context, contract client.Contract) ([]*domain.OpenRevUser, error) {
	log.Println("Evaluate Transaction: ReadAllOpenRevUserAssets, function returns all the users assets on the ledger")

	evaluateResult, err := contract.EvaluateTransaction("ReadAllOpenRevUserAssets")
	if err != nil {
		return nil, helper.LedgerErrorHandler(&contract, err)
	}

	var users []*domain.OpenRevUser

	err = json.Unmarshal(evaluateResult, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userUsecase) Register(context context.Context, contract client.Contract, user domain.OpenRevUser) (*domain.OpenRevUser, error) {
	user.ID = uuid.New().String()
	user.RoleId = 4
	user.Verified = true
	code := helper.RandomStringGenerator(8)

	log.Println("Submit Transaction: CreateRevUserAsset, function creates user on the ledger")

	evaluateResult, err := contract.SubmitTransaction("CreateRevUserAsset", user.ID, user.Name, user.Surname, code, user.Email)
	if err != nil {
		return nil, helper.LedgerErrorHandler(&contract, err)
	}

	var openRevUser domain.OpenRevUser
	err = json.Unmarshal(evaluateResult, &openRevUser)
	if err != nil {
		return nil, err
	}

	return &openRevUser, nil
}

func (r *userUsecase) ConfirmAccount(context context.Context, contract client.Contract, code, id string) (*domain.OpenRevUser, error) {
	log.Printf("Submit Transaction: VerifyRevUserAsset, function verifies the user with id %s", id)

	evaluateResult, err := contract.SubmitTransaction("VerifyRevUserAsset", code, id)
	if err != nil {
		return nil, helper.LedgerErrorHandler(&contract, err)
	}

	var openRevUser domain.OpenRevUser
	err = json.Unmarshal(evaluateResult, &openRevUser)
	if err != nil {
		return nil, err
	}

	return &openRevUser, nil
}

type UserUsecase interface {
	EditUser(context context.Context, contract client.Contract, user dto.EditUserDTO) error
	GetAllUsers(context context.Context, contract client.Contract) ([]*domain.OpenRevUser, error)
	Register(context context.Context, contract client.Contract, user domain.OpenRevUser) (*domain.OpenRevUser, error)
	ConfirmAccount(context context.Context, contract client.Contract, code, id string) (*domain.OpenRevUser, error)
	GetUserById(context context.Context, contract client.Contract, id string) (*domain.OpenRevUser, error)
	GetUserInfo(context context.Context, contract client.Contract, id string) (*dto.OpenRevUserInfoDTO, error)
	GetAllUsersWithDetails(context context.Context, contract client.Contract) ([]*dto.OpenRevUserInfoDTO, error)

	DeleteOpenRevUser(context context.Context, contract client.Contract, id string) error
}

func NewUserUsecase(contract *client.Contract) UserUsecase {
	return &userUsecase{Contract: contract}
}
