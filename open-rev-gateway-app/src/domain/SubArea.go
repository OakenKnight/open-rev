package domain

import "time"

type SubArea struct {
	ID             string
	Name           string
	AreaId         string
	Hidden         bool
	IsDeleted      bool
	LastUpdateTime time.Time `json:"last_update_time"`
}
