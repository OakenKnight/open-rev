package dto

type AddAreaDto struct {
	Name string `json:"area_name"`
}

type AddSubAreaDto struct {
	Name   string `json:"area_name"`
	AreaId string `json:"area_id"`
}
