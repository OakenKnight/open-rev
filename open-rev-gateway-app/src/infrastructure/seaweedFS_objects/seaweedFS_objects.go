package seaweedFS_objects

type FidResponse struct {
	FID       string `json:"fid"`
	Url       string `json:"url"`
	PublicUrl string `json:"publicUrl"`
	Count     int    `json:"count"`
}

type UploadAnswer struct {
	Name string `json:"name"`
	Size int    `json:"size"`
	ETag string `json:"eTag"`
}
