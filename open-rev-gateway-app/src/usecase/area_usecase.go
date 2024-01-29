package usecase

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"log"
	"open-rev.com/helper"
	"open-rev.com/infrastructure/dto"
	"strconv"
)

type areaUsecase struct {
	Contract *client.Contract
}

func (a *areaUsecase) AddArea(context context.Context, contract client.Contract, areaDto dto.AddAreaDto) error {
	log.Println("Submit Transaction: CreateAreaAsset, function returns error if not successful")
	ID := uuid.New().String()
	time := strconv.FormatInt(areaDto.LastUpdateTime.UnixMilli(), 10)
	_, err := contract.SubmitTransaction("CreateAreaAsset", ID, areaDto.Name, time)
	if err != nil {
		return helper.LedgerErrorHandler(&contract, err)
	}
	return nil
}
func (a *areaUsecase) AddSubArea(context context.Context, contract client.Contract, areaDto dto.AddSubAreaDto) error {
	log.Println("Submit Transaction: CreateSubAreaAsset, function returns error if not successful")
	ID := uuid.New().String()
	time := strconv.FormatInt(areaDto.LastUpdateTime.UnixMilli(), 10)
	_, err := contract.SubmitTransaction("CreateSubAreaAsset", ID, areaDto.Name, areaDto.AreaId, time)
	if err != nil {
		return helper.LedgerErrorHandler(&contract, err)
	}
	return nil
}
func (a *areaUsecase) GetAllSubAreas(context context.Context, contract client.Contract) ([]*dto.SubAreaDto, error) {
	log.Println("Evaluate Transaction: ReadAllSubAreaAssets, function returns all the subArea assets on the ledger")

	evaluateResult, err := contract.EvaluateTransaction("ReadAllSubAreaAssets")
	if err != nil {
		return nil, helper.LedgerErrorHandler(&contract, err)
	}
	var subareas []*dto.SubAreaDto
	err = json.Unmarshal(evaluateResult, &subareas)
	if err != nil {
		return nil, err
	}

	return subareas, nil
}
func (a *areaUsecase) DeleteArea(context context.Context, contract client.Contract, id string) error {
	log.Println("Submit Transaction: DeleteAreaAsset, function returns error if not successful")
	_, err := contract.SubmitTransaction("DeleteAreaAsset", id)
	if err != nil {
		return helper.LedgerErrorHandler(&contract, err)
	}
	return nil
}
func (a *areaUsecase) DeleteSubArea(context context.Context, contract client.Contract, id string) error {
	log.Println("Submit Transaction: DeleteSubArea, function returns error if not successful")
	_, err := contract.SubmitTransaction("DeleteSubAreaAsset", id)
	if err != nil {
		return helper.LedgerErrorHandler(&contract, err)
	}
	return nil
}
func (a *areaUsecase) GetAllAreas(context context.Context, contract client.Contract) ([]*dto.AreaDto, error) {
	log.Println("Evaluate Transaction: ReadAllAreaAssets, function returns all area assets on the ledger")

	evaluateResult, err := contract.EvaluateTransaction("ReadAllAreaAssets")
	if err != nil {
		return nil, helper.LedgerErrorHandler(&contract, err)
	}

	var areas []*dto.AreaDto

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
	GetAllAreas(context context.Context, contract client.Contract) ([]*dto.AreaDto, error)
	GetAllSubAreas(context context.Context, contract client.Contract) ([]*dto.SubAreaDto, error)
	GetAllAreasAndSubAreas(context context.Context, contract client.Contract) ([]*dto.AreaSubareaDTO, error)
	DeleteArea(context context.Context, contract client.Contract, id string) error
	DeleteSubArea(context context.Context, contract client.Contract, id string) error
	AddArea(context context.Context, contract client.Contract, areaDto dto.AddAreaDto) error
	AddSubArea(context context.Context, contract client.Contract, areaDto dto.AddSubAreaDto) error
}

func NewAreaUsecase(contract *client.Contract) AreaUsecase {
	return &areaUsecase{Contract: contract}
}
