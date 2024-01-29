package domain

type Review struct {
	ID               string
	Review           string
	Assessment       int
	Recommend        bool
	UserId           string
	ScientificWorkId string
	IsDeleted        bool
	LastUpdateTime   string `json:"last_update_time"`
}
