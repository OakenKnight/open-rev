package domain

type Area struct {
	ID     string
	Name   string
	Hidden bool
	Type   string
	IsDeleted   bool

}
type SubArea struct {
	ID     string
	Name   string
	AreaId string
	Hidden bool
	Type   string
	IsDeleted   bool

}
type Role struct {
	ID   string
	Name string
	Type string
	IsDeleted   bool

}
type OpenRevUser struct {
	ID       string
	Name     string
	Surname  string
	Email    string
	RoleId   int
	Verified bool
	Code     string
	Type     string
	IsDeleted   bool

}

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
	Type        string `json:"type"`
}

type DashboardItem struct {
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
	IsDeleted bool `json:"is_deleted"`
}

type AreaSubareaDTO struct {
	Area     string       `json:"area"`
	SubAreas []SubAreaDTO `json:"subareas"`
}

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
	IsDeleted bool `json:"is_deleted"`
}

type Review struct {
	ID               string
	Review           string
	Assessment       int
	Recommend        bool
	UserId           string
	ScientificWorkId string
	Type             string
	IsDeleted   bool

}

type ReviewQuality struct {
	ID         string
	Assessment int
	UserId     string
	ReviewId   string
	Type       string
	IsDeleted   bool
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

type ScientificWorkDetailsDTO struct {
	Area     string                `json:"area"`
	AvgMark  float32               `json:"avg_mark"`
	WorkInfo ScientificWork 	   `json:"work_info"`
	Review   []ReviewForDetailsDTO `json:"review"`
}


type ScientificWorkForSortingDTO struct {
	ID string `json:"guid"`
	Title string `json:"title"`
	PublishDate string `json:"publish_date"`
	Abstract string `json:"abstract"`
	Keywords string `json:"keywords"`
	PdfFile string `json:"pdf"`
	User string `json:"user"`
	AvgRate float32 `json:"avg_rate"`
}