package dto

type ReviewDTO struct {
	ID               string
	ScientificWorkId string
	UserId           string
	Assessment       int
	Recommend        bool
	Review           string
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

//
