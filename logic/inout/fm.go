package inout

import "mime/multipart"

type FM struct {
	Path  string `json:"path"  query:"path"`
	Match string `json:"match" query:"match"`
	Limit int    `json:"limit" query:"limit"`
}

type AddFile struct {
	Name string                `form:"name" validate:"required,lte=100"`
	Path string                `form:"path" validate:"required"`
	File *multipart.FileHeader `form:"file" validate:"required"`
}
