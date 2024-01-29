package dto

import "time"

type ReviewQualityDTO struct {
	ID             string    `json:"id"`
	Assessment     int       `json:"assessment"`
	UserId         string    `json:"userId"`
	ReviewId       string    `json:"reviewId"`
	LastUpdateTime time.Time `json:"last_update_time"`
}
