package dto

type ReviewQualityDTO struct {
	ID         string `json:"id"`
	Assessment int    `json:"assessment"`
	UserId     string `json:"userId"`
	ReviewId   string `json:"reviewId"`
}
