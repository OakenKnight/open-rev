package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"log"
	"open-rev.com/domain"
	"open-rev.com/helper"
	"open-rev.com/infrastructure/dto"
	"sort"
	"strconv"
	"time"
)

type scientificWorkUsecase struct {
	Contract *client.Contract
}

func (s *scientificWorkUsecase) GetAllScientificWorksWithDetails(context context.Context, contract client.Contract) ([]*dto.ScientificWorkWithDetailsDTO, error) {
	log.Printf("Evaluate Transaction: ReadAllScientificWorkAssetsWithDetails, function returns scientific works with details from the ledger")

	evaluateResult, err := contract.EvaluateTransaction("ReadAllScientificWorkAssetsWithDetails")
	if err != nil {
		return nil, helper.LedgerErrorHandler(&contract, err)
	}

	var works []*dto.ScientificWorkWithDetailsDTO

	err = json.Unmarshal(evaluateResult, &works)
	if err != nil {
		return nil, err
	}
	for _, work := range works {
		publishDate, err := convertStringToDate(work.PublishDate)
		if err != nil {
			return nil, err
		}
		work.PublishDate = publishDate.String()

	}
	return works, nil
}

func (s *scientificWorkUsecase) DeleteSciWork(context context.Context, contract client.Contract, id string) error {
	log.Println("Submit Transaction: DeleteScientificWorkAsset, function returns error if not successful")
	_, err := contract.SubmitTransaction("DeleteScientificWorkAsset", id)
	if err != nil {
		return helper.LedgerErrorHandler(&contract, err)
	}
	return nil
}

func (s *scientificWorkUsecase) GetAllScientificWorksBySubareaId(context context.Context, contract client.Contract, id string) ([]*domain.ScientificWork, error) {
	log.Printf("Evaluate Transaction: ReadAllScientificWorksBySubAreaAssets, function returns all the scientific works that belong to subarea with id %s", id)

	evaluateResult, err := contract.EvaluateTransaction("ReadAllScientificWorksBySubAreaAssets", id)
	if err != nil {
		return nil, helper.LedgerErrorHandler(&contract, err)
	}

	if evaluateResult == nil {
		return make([]*domain.ScientificWork, 0), nil
	}

	var works []*domain.ScientificWork

	err = json.Unmarshal(evaluateResult, &works)
	if err != nil {
		return nil, err
	}
	for _, work := range works {
		publishDate, err := convertStringToDate(work.PublishDate)
		if err != nil {
			return nil, err
		}
		work.PublishDate = publishDate.String()
	}
	return works, nil
}

func (s *scientificWorkUsecase) GetAllScientificWorksByUser(context context.Context, contract client.Contract, userId string) ([]*domain.ScientificWork, error) {
	log.Printf("Evaluate Transaction: ReadAllScientificWorksByUserAssets, function returns all scientific works by user %s", userId)
	evaluateResult, err := contract.EvaluateTransaction("ReadAllScientificWorksByUserAssets", userId)
	if err != nil {
		return nil, helper.LedgerErrorHandler(&contract, err)
	}

	var works []*domain.ScientificWork

	err = json.Unmarshal(evaluateResult, &works)
	if err != nil {
		return nil, err
	}
	for _, work := range works {
		publishDate, err := convertStringToDate(work.PublishDate)
		if err != nil {
			return nil, err
		}
		work.PublishDate = publishDate.String()
	}
	return works, nil
}

func (s *scientificWorkUsecase) GetScientificWorkById(context context.Context, contract client.Contract, id string) (*domain.ScientificWork, error) {
	log.Printf("Evaluate Transaction: ReadScientificWorkAsset, function returns scientific work with id %s", id)

	evaluateResult, err := contract.EvaluateTransaction("ReadScientificWorkAsset", id)
	if err != nil {
		return nil, helper.LedgerErrorHandler(&contract, err)
	}

	var work *domain.ScientificWork

	err = json.Unmarshal(evaluateResult, &work)
	if err != nil {
		return nil, err
	}
	publishDate, err := convertStringToDate(work.PublishDate)
	if err != nil {
		return nil, err
	}
	work.PublishDate = publishDate.String()
	return work, nil
}

func (s *scientificWorkUsecase) GetScientificWorkDetails(context context.Context, contract client.Contract, id string) (*dto.ScientificWorkDetailsDTO, error) {
	log.Printf("Evaluate Transaction: ReadScientificWorkDetails, function returns details of scientific work with id %s on the ledger", id)

	evaluateResult, err := contract.EvaluateTransaction("ReadScientificWorkDetails", id)
	if err != nil {
		return nil, helper.LedgerErrorHandler(&contract, err)
	}

	var work *dto.ScientificWorkDetailsDTO

	err = json.Unmarshal(evaluateResult, &work)
	if err != nil {
		return nil, err
	}
	publishDate, err := convertStringToDate(work.WorkInfo.PublishDate)
	if err != nil {
		return nil, err
	}
	work.WorkInfo.PublishDate = publishDate.String()
	return work, nil
}
func (s *scientificWorkUsecase) CreateScientificWork(context context.Context, contract client.Contract, dto *dto.NewScientificWorkDTO) (*domain.ScientificWork, error) {
	sciWork := domain.ScientificWork{Keywords: dto.Keywords, Abstract: dto.Abstract, UserId: dto.UserId, SubAreaId: dto.SubAreaId, Title: dto.Title, PdfFile: dto.PdfFile}
	sciWork.ID = uuid.New().String()

	time := strconv.FormatInt(time.Now().UnixMilli(), 10)

	log.Println("Submit Transaction: CreateScientificWorkAsset, function creates scientific work on the ledger")
	evaluateResult, err := contract.SubmitTransaction("CreateScientificWorkAsset", sciWork.ID, sciWork.Title, sciWork.Abstract, sciWork.Keywords, sciWork.PdfFile, sciWork.SubAreaId, sciWork.UserId, time)
	if err != nil {
		return nil, helper.LedgerErrorHandler(&contract, err)
	}

	var scientificWork *domain.ScientificWork
	err = json.Unmarshal(evaluateResult, &scientificWork)
	if err != nil {
		return nil, err
	}

	return scientificWork, nil
}

func (s *scientificWorkUsecase) GetAllScientificWorks(context context.Context, contract client.Contract) ([]*domain.ScientificWork, error) {
	log.Println("Evaluate Transaction: ReadAllScientificWorkAssets, function returns all scientific works on the ledger")

	evaluateResult, err := contract.EvaluateTransaction("ReadAllScientificWorkAssets")
	if err != nil {
		return nil, helper.LedgerErrorHandler(&contract, err)
	}

	var works []*domain.ScientificWork

	err = json.Unmarshal(evaluateResult, &works)
	if err != nil {
		return nil, err
	}
	for _, work := range works {
		publishDate, err := convertStringToDate(work.PublishDate)
		if err != nil {
			return nil, err
		}
		work.PublishDate = publishDate.String()
	}
	return works, nil
}

func (s *scientificWorkUsecase) GetDashboard(context context.Context, contract client.Contract) (*dto.DashboardDTO, error) {
	log.Println("Evaluate Transaction: ReadAllDashboardItemAssets, function returns all items for the dashboard on the ledger")

	evaluateResult, err := contract.EvaluateTransaction("ReadAllDashboardItemAssets")
	if err != nil {
		return nil, helper.LedgerErrorHandler(&contract, err)
	}

	var works []*dto.DashboardItemDTO

	err = json.Unmarshal(evaluateResult, &works)
	if err != nil {
		return nil, err
	}
	dashboard := dto.DashboardDTO{
		MostRecent:  make([]dto.DashboardItemForSortDTO, 0),
		Assessments: make([]dto.DashboardItemForSortDTO, 0),
	}

	sort.Slice(works, func(i, j int) bool {
		return works[i].AverageRate > works[j].AverageRate
	})

	for _, work := range works {

		publishDate, err := convertStringToDate(work.PublishDate)
		if err != nil {
			return nil, err
		}
		lastUpdateTime, err := convertStringToDate(work.LastUpdateTime)
		if err != nil {
			return nil, err
		}

		dashboard.Assessments = append(dashboard.Assessments, dto.DashboardItemForSortDTO{
			ID:             work.ID,
			User:           work.User,
			AverageRate:    work.AverageRate,
			Abstract:       work.Abstract,
			Keywords:       work.Keywords,
			PdfFile:        work.PdfFile,
			Title:          work.Title,
			PublishDate:    publishDate,
			LastUpdateTime: lastUpdateTime,
		})
	}

	for _, v := range dashboard.Assessments {
		dashboard.MostRecent = append(dashboard.MostRecent, v)
	}
	sort.Slice(dashboard.MostRecent, func(i, j int) bool {
		return dashboard.MostRecent[i].LastUpdateTime.After(dashboard.MostRecent[j].PublishDate)
	})

	return &dashboard, nil
}
func convertStringToDate(timestamp string) (time.Time, error) {
	milliseconds, err := strconv.Atoi(timestamp)
	if err != nil {
		return time.Now(), fmt.Errorf("Error parsing date")
	}
	return time.UnixMilli(int64(milliseconds)), nil
}

type ScientificWorkUsecase interface {
	GetAllScientificWorks(context context.Context, contract client.Contract) ([]*domain.ScientificWork, error)
	GetAllScientificWorksWithDetails(context context.Context, contract client.Contract) ([]*dto.ScientificWorkWithDetailsDTO, error)
	GetAllScientificWorksByUser(context context.Context, contract client.Contract, userId string) ([]*domain.ScientificWork, error)
	GetScientificWorkById(context context.Context, contract client.Contract, id string) (*domain.ScientificWork, error)
	GetScientificWorkDetails(context context.Context, contract client.Contract, id string) (*dto.ScientificWorkDetailsDTO, error)
	CreateScientificWork(context context.Context, contract client.Contract, sciWork *dto.NewScientificWorkDTO) (*domain.ScientificWork, error)
	GetDashboard(context context.Context, contract client.Contract) (*dto.DashboardDTO, error)
	GetAllScientificWorksBySubareaId(context context.Context, contract client.Contract, id string) ([]*domain.ScientificWork, error)
	DeleteSciWork(context context.Context, contract client.Contract, id string) error
}

func NewScientificWorkUsecase(contract *client.Contract) ScientificWorkUsecase {
	return &scientificWorkUsecase{Contract: contract}
}
