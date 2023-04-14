package dto

import "time"

type DashboardItemDTO struct {
	User        string  `json:"user"`
	Title       string  `json:"title"`
	PublishDate string  `json:"publish_date"`
	AverageRate float32 `json:"avg_rate"`
	Abstract    string  `json:"abstract"`
	Keywords    string  `json:"keywords"`
	PdfFile     string  `json:"pdf"`
	ID          string  `json:"guid"`
}

type SubAreaDTO struct {
	ID      string `json:"id"`
	SubArea string `json:"subarea"`
}

type AreaSubareaDTO struct {
	Area     string       `json:"area"`
	SubAreas []SubAreaDTO `json:"subareas"`
}

type DashboardItemForSortDTO struct {
	User        string    `json:"user"`
	Title       string    `json:"title"`
	PublishDate time.Time `json:"publish_date"`
	AverageRate float32   `json:"avg_rate"`
	Abstract    string    `json:"abstract"`
	Keywords    string    `json:"keywords"`
	PdfFile     string    `json:"pdf"`
	ID          string    `json:"guid"`
}

type DashboardDTO struct {
	MostRecent  []DashboardItemForSortDTO `json:"most_recent"`
	Assessments []DashboardItemForSortDTO `json:"assessments"`
}
