package v1

import "mime/multipart"

// ImageUpload ..
type ImageUpload struct {
	FilePath string                `json:"filepath" form:"filepath" binding:"required"`
	File     *multipart.FileHeader `json:"file" form:"file" binding:"required"`
}
