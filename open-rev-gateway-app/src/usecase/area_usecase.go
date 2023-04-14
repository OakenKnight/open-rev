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

type areaUsecase struct {
	Contract *client.Contract
}

func (a *areaUsecase) AddArea(context context.Context, contract client.Contract, areaDto dto.AddAreaDto) error {
	log.Println("Evaluate Transaction: CreateAreaAsset, function returns error if not successful")
	ID := uuid.New().String()

	_, err := contract.SubmitTransaction("CreateAreaAsset", ID, areaDto.Name)
	if err != nil {
		return helper.LedgerErrorHandler(&contract, err)
	}
	return nil
}
func (a *areaUsecase) AddSubArea(context context.Context, contract client.Contract, areaDto dto.AddSubAreaDto) error {
	log.Println("Evaluate Transaction: CreateSubAreaAsset, function returns error if not successful")
	ID := uuid.New().String()

	_, err := contract.SubmitTransaction("CreateSubAreaAsset", ID, areaDto.Name, areaDto.AreaId)
	if err != nil {
		return helper.LedgerErrorHandler(&contract, err)
	}
	return nil
}
func (a *areaUsecase) GetAllSubAreas(context context.Context, contract client.Contract) ([]*domain.SubArea, error) {
	log.Println("Evaluate Transaction: ReadAllSubAreaAssets, function returns all the subArea assets on the ledger")

	evaluateResult, err := contract.EvaluateTransaction("ReadAllSubAreaAssets")
	if err != nil {
		return nil, helper.LedgerErrorHandler(&contract, err)
	}
	var subareas []*domain.SubArea
	err = json.Unmarshal(evaluateResult, &subareas)
	if err != nil {
		return nil, err
	}

	return subareas, nil
}
func (a *areaUsecase) DeleteArea(context context.Context, contract client.Contract, id string) error {
	log.Println("Evaluate Transaction: DeleteAreaAsset, function returns error if not successful")
	_, err := contract.SubmitTransaction("DeleteAreaAsset", id)
	if err != nil {
		return helper.LedgerErrorHandler(&contract, err)
	}
	return nil
}
func (a *areaUsecase) DeleteSubArea(context context.Context, contract client.Contract, id string) error {
	log.Println("Evaluate Transaction: DeleteSubArea, function returns error if not successful")
	_, err := contract.SubmitTransaction("DeleteSubAreaAsset", id)
	if err != nil {
		return helper.LedgerErrorHandler(&contract, err)
	}
	return nil
}
func (a *areaUsecase) GetAllAreas(context context.Context, contract client.Contract) ([]*domain.Area, error) {
	log.Println("Evaluate Transaction: ReadAllAreaAssets, function returns all area assets on the ledger")

	evaluateResult, err := contract.EvaluateTransaction("ReadAllAreaAssets")
	if err != nil {
		return nil, helper.LedgerErrorHandler(&contract, err)
	}


	var areas []*domain.Area


	err = json.Unmarshal(evaluateResult, &areas)
	if err != nil {
		return nil, err
	}



	return areas, nil
}

func (a *areaUsecase) GetAllAreasAndSubAreas(context context.Context, contract client.Contract) ([]*dto.AreaSubareaDTO, error) {
	log.Println("Evaluate Transaction: ReadAllAreaSubareaAssets, function returns all the area with subarea assets on the ledger")

	evaluateResult, err := contract.EvaluateTransaction("ReadAllAreaSubareaAssets")
	if err != nil {
		return nil, helper.LedgerErrorHandler(&contract, err)
	}

	var areas []*dto.AreaSubareaDTO

	err = json.Unmarshal(evaluateResult, &areas)
	if err != nil {
		return nil, err
	}

	return areas, nil
}

type AreaUsecase interface {
	GetAllAreas(context context.Context, contract client.Contract) ([]*domain.Area, error)
	GetAllSubAreas(context context.Context, contract client.Contract) ([]*domain.SubArea, error)
	GetAllAreasAndSubAreas(context context.Context, contract client.Contract) ([]*dto.AreaSubareaDTO, error)
	DeleteArea(context context.Context, contract client.Contract, id string) error
	DeleteSubArea(context context.Context, contract client.Contract, id string) error
	AddArea(context context.Context, contract client.Contract, areaDto dto.AddAreaDto) error
	AddSubArea(context context.Context, contract client.Contract, areaDto dto.AddSubAreaDto) error
}

func NewAreaUsecase(contract *client.Contract) AreaUsecase {
	return &areaUsecase{Contract: contract}
}
