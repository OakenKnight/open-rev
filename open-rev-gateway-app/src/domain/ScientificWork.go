package domain

type ScientificWork struct {
	ID          string `json:"guid"`
	SubAreaId   string `json:"sub_area_id"`
	Title       string `json:"title"`
	PublishDate string `json:"publish_date"`
	Abstract    string `json:"abstract"`
	Keywords    string `json:"keywords"`
	PdfFile     string `json:"pdf"`
	UserId      string `json:"user_guid"`
	IsDeleted   bool   `json:"is_deleted"`
}

type ScientificWorkDetailsDTO struct {
	Area     string
	AvgMarg  float32
	WorkInfo ScientificWork
}
