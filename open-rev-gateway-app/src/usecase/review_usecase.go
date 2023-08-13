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
	"strconv"
)

type reviewUsecase struct {
	Contract *client.Contract
}

func (r *reviewUsecase) GetAllReviewQualities(ctx context.Context, contract client.Contract) ([]*domain.ReviewQuality, error) {
	log.Println("Evaluate Transaction: ReadAllReviewQualityAssets, function returns all the reviews on the ledger")

	evaluateResult, err := contract.EvaluateTransaction("ReadAllReviewQualityAssets")
	if err != nil {
		return nil, helper.LedgerErrorHandler(&contract, err)
	}

	var reviewsQualities []*domain.ReviewQuality

	err = json.Unmarshal(evaluateResult, &reviewsQualities)
	if err != nil {
		return nil, err
	}

	return reviewsQualities, nil
}

func (r *reviewUsecase) FixReviewId(context context.Context, contract client.Contract, id string) error {
	log.Println("Submit Transaction: DeleteOffReview, function returns error if not successful")
	_, err := contract.SubmitTransaction("DeleteOffReview", id)
	if err != nil {
		return helper.LedgerErrorHandler(&contract, err)
	}
	return nil
}

func (r *reviewUsecase) FixReviewQualityId(context context.Context, contract client.Contract, id string) error {
	log.Println("Submit Transaction: DeleteOffReviewQuality, function returns error if not successful")
	_, err := contract.SubmitTransaction("DeleteOffReviewQuality", id)
	if err != nil {
		return helper.LedgerErrorHandler(&contract, err)
	}
	return nil
}

func (r *reviewUsecase) GetAllReviewQualityByReview(ctx context.Context, contract client.Contract, reviewId string) ([]*domain.ReviewQuality, error) {
	log.Printf("Evaluate Transaction: ReadAllReviewsQualityByReviewAssets, function returns all reviews for scientific work %s on the ledger", reviewId)

	evaluateResult, err := contract.EvaluateTransaction("ReadAllReviewsQualityByReviewAssets", reviewId)

	if err != nil {
		return nil, helper.LedgerErrorHandler(&contract, err)
	}

	var reviews []*domain.ReviewQuality
	if evaluateResult == nil {
		return make([]*domain.ReviewQuality, 0), nil
	}
	err = json.Unmarshal(evaluateResult, &reviews)
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *reviewUsecase) DeleteReview(context context.Context, contract client.Contract, id string) error {
	log.Println("Submit Transaction: DeleteReview, function returns error if not successful")
	_, err := contract.SubmitTransaction("DeleteReviewAsset", id)
	if err != nil {
		return helper.LedgerErrorHandler(&contract, err)
	}
	return nil
}

func (r *reviewUsecase) CreateReviewQuality(ctx context.Context, contract client.Contract, revQ *dto.ReviewQualityDTO) (*domain.ReviewQuality, error) {

	revQ.ID = uuid.New().String()

	log.Printf("Submit Transaction: CreateReviewQualityAsset, function creates ReviewQuality for review %s asset on the ledger", revQ.ReviewId)
	evaluateResult, err := contract.SubmitTransaction("CreateReviewQualityAsset", revQ.ID, revQ.ReviewId, revQ.UserId, strconv.Itoa(revQ.Assessment))
	if err != nil {
		return nil, helper.LedgerErrorHandler(&contract, err)
	}

	var reviewAsset *domain.ReviewQuality
	err = json.Unmarshal(evaluateResult, &reviewAsset)
	if err != nil {
		return nil, err
	}

	return reviewAsset, nil
}

func (r *reviewUsecase) CreateReview(ctx context.Context, contract client.Contract, review dto.ReviewDTO) (*domain.Review, error) {

	review.ID = uuid.New().String()

	log.Printf("Submit Transaction: CreateReviewAsset, function creates review for scientific work %s on the ledger", review.ScientificWorkId)

	evaluateResult, err := contract.SubmitTransaction("CreateReviewAsset", review.ID, review.ScientificWorkId, review.UserId, strconv.Itoa(review.Assessment), strconv.FormatBool(review.Recommend), review.Review)

	if err != nil {
		return nil, helper.LedgerErrorHandler(&contract, err)
	}

	var reviewAsset *domain.Review
	err = json.Unmarshal(evaluateResult, &reviewAsset)
	if err != nil {
		return nil, err
	}

	return reviewAsset, nil
}

func (r *reviewUsecase) GetAllReviewsByScientificWork(ctx context.Context, contract client.Contract, sciWorkId string) ([]*domain.Review, error) {
	log.Printf("Evaluate Transaction: ReadAllReviewsByScientificPaperAssets, function returns all reviews for scientific work %s on the ledger", sciWorkId)

	evaluateResult, err := contract.EvaluateTransaction("ReadAllReviewsByScientificPaperAssets", sciWorkId)

	if err != nil {
		return nil, helper.LedgerErrorHandler(&contract, err)
	}

	var reviews []*domain.Review
	if evaluateResult == nil {
		return make([]*domain.Review, 0), nil
	}
	err = json.Unmarshal(evaluateResult, &reviews)
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *reviewUsecase) GetAllReviewsByUser(ctx context.Context, contract client.Contract, sciWorkId string) ([]*domain.Review, error) {
	log.Printf("Evaluate Transaction: ReadAllReviewsByScientificPaperAssets, function returns all reviews for scientific work %s on the ledger", sciWorkId)

	evaluateResult, err := contract.EvaluateTransaction("ReadAllReviewsByOpenRevUserAssets", sciWorkId)

	if err != nil {
		return nil, helper.LedgerErrorHandler(&contract, err)
	}

	var reviews []*domain.Review

	err = json.Unmarshal(evaluateResult, &reviews)
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *reviewUsecase) GetAllReviews(context context.Context, contract client.Contract) ([]*domain.Review, error) {
	log.Println("Evaluate Transaction: ReadAllReviewAssets, function returns all the reviews on the ledger")

	evaluateResult, err := contract.EvaluateTransaction("ReadAllReviewAssets")
	if err != nil {
		return nil, helper.LedgerErrorHandler(&contract, err)
	}

	var reviews []*domain.Review

	err = json.Unmarshal(evaluateResult, &reviews)
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

type ReviewUsecase interface {
	GetAllReviews(ctx context.Context, contract client.Contract) ([]*domain.Review, error)
	GetAllReviewsByScientificWork(ctx context.Context, contract client.Contract, sciWorkId string) ([]*domain.Review, error)
	GetAllReviewsByUser(ctx context.Context, contract client.Contract, sciWorkId string) ([]*domain.Review, error)
	CreateReview(ctx context.Context, contract client.Contract, review dto.ReviewDTO) (*domain.Review, error)
	CreateReviewQuality(ctx context.Context, contract client.Contract, review *dto.ReviewQualityDTO) (*domain.ReviewQuality, error)
	DeleteReview(context context.Context, contract client.Contract, id string) error
	FixReviewId(context context.Context, contract client.Contract, id string) error
	FixReviewQualityId(context context.Context, contract client.Contract, id string) error
	GetAllReviewQualities(ctx context.Context, contract client.Contract) ([]*domain.ReviewQuality, error)
	GetAllReviewQualityByReview(ctx context.Context, contract client.Contract, reviewId string) ([]*domain.ReviewQuality, error)
}

func NewReviewUsecase(contract *client.Contract) ReviewUsecase {
	return &reviewUsecase{Contract: contract}

}
