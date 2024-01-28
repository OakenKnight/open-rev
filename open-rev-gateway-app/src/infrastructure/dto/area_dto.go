package dto

import "time"

type AddAreaDto struct {
	Name           string    `json:"area_name"`
	LastUpdateTime time.Time `json:"last_update_time"`
}

type AddSubAreaDto struct {
	Name           string    `json:"area_name"`
	AreaId         string    `json:"area_id"`
	LastUpdateTime time.Time `json:"last_update_time"`
}
