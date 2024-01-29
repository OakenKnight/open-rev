package dto

import (
	"open-rev.com/domain"
	"time"
)

type ScientificWorkDetailsDTO struct {
	Area     string                `json:"area"`
	AvgMark  float32               `json:"avg_mark"`
	WorkInfo domain.ScientificWork `json:"work_info"`
	Review   []ReviewForDetailsDTO `json:"review"`
}

type NewScientificWorkDTO struct {
	SubAreaId      string    `json:"SubAreaId"`
	Title          string    `json:"Title"`
	Abstract       string    `json:"Abstract"`
	Keywords       string    `json:"Keywords"`
	PdfFile        string    `json:"PdfFile"`
	UserId         string    `json:"UserId"`
	LastUpdateTime time.Time `json:"last_update_time"`
}

type ScientificWorkWithDetailsDTO struct {
	ID             string  `json:"guid"`
	Title          string  `json:"title"`
	SubAreaId      string  `json:"sub_area_id"`
	PublishDate    string  `json:"publish_date"`
	Abstract       string  `json:"abstract"`
	Keywords       string  `json:"keywords"`
	PdfFile        string  `json:"pdf"`
	User           string  `json:"user"`
	AvgRate        float32 `json:"avg_rate"`
	LastUpdateTime string  `json:"last_update_time"`
}
