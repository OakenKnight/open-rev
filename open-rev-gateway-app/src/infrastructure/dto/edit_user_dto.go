package dto

type EditUserDTO struct {
	Name    string `json:"Name"`
	Surname string `json:"Surname"`
	ID      string `json:"ID"`
}

type ChangePasswordsDTO struct {
	OldPassword  string `json:"OldPassword"`
	NewPassword1 string `json:"NewPassword1"`
	NewPassword2 string `json:"NewPassword2"`
	ID           string `json:"ID"`
}
