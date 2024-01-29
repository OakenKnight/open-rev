package domain

type ReviewQuality struct {
	ID             string
	Assessment     int
	UserId         string
	ReviewId       string
	IsDeleted      bool
	LastUpdateTime string `json:"last_update_time"`
}
