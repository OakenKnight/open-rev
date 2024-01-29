package dto

import "time"

type EditUserDTO struct {
	Name           string    `json:"Name"`
	Surname        string    `json:"Surname"`
	ID             string    `json:"ID"`
	LastUpdateTime time.Time `json:"last_update_time"`
}
