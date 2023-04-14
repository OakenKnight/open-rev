package domain

type ReviewQuality struct {
	ID         string
	Assessment int
	UserId     string
	ReviewId   string
	IsDeleted  bool
}
