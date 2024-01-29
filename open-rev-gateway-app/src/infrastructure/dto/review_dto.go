package dto

type ReviewDTO struct {
	ID               string
	ScientificWorkId string
	UserId           string
	Assessment       int
	Recommend        bool
	Review           string
	LastUpdateTime   string `json:"last_update_time"`
}
type NewReviewDTO struct {
	ScientificWorkId string `json:"scientificWorkId"`
	UserId           string `json:"userId"`
	Assessment       int    `json:"assessment"`
	Recommend        bool   `json:"recommend"`
	Review           string `json:"review"`
}
type ReviewForDetailsDTO struct {
	ReviewId         string   `json:"review_id"`
	Review           string   `json:"review"`
	Assessment       string   `json:"assessment"`
	CountRevQuality  int      `json:"count_rev_quality"`
	SumRevQuality    int      `json:"sum_rev_quality"`
	Recommend        string   `json:"recommend"`
	UserId           string   `json:"user_id"`
	User             string   `json:"user"`
	ScientificWorkId string   `json:"scientific_work_id"`
	UsersRevQuality  []string `json:"users_rev_quality"`
}
