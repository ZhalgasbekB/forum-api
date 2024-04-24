package model

type UploadPostImage struct {
	Path   string `json:"path"`
	PostId int    `json:"post_id"`
}
