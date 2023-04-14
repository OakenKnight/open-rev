package domain

type OpenRevUser struct {
	ID        string
	Name      string
	Surname   string
	Email     string
	RoleId    int
	Code      string
	Verified  bool
	IsDeleted bool
}
