package dto

type OpenRevUserInfoDTO struct {
	ID               string  `json:"guid"`
	Name             string  `json:"name"`
	Surname          string  `json:"surname"`
	AvgMark          float32 `json:"avg_mark"`
	AvgMyRevsQuality float32 `json:"avg_my_revs_quality"`
	AvgRevQuality    float32 `json:"avg_rev_quality"`
	AvgReview        float32 `json:"avg_review"`
	Email            string  `json:"email"`
	RoleId           int     `json:"role"`
	WorksCount       int     `json:"works_count"`
	ReviewsCount     int     `json:"reviews_count"`
	IsAdmin          bool    `json:"isAdmin"`
}

type TopAuthorDto struct {
	User    string  `json:"user"`
	AvgRate float32 `json:"avg_rate"`
	Guid    string  `json:"guid"`
}

type TopReviewerDto struct {
	User      string  `json:"user"`
	AvgReview float32 `json:"avg_review"`
	Guid      string  `json:"guid"`
}
