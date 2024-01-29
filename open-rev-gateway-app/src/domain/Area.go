package domain

import "time"

type Area struct {
	ID             string
	Name           string
	Hidden         bool
	IsDeleted      bool
	LastUpdateTime time.Time `json:"last_update_time"`
}
