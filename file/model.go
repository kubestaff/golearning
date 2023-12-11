package file

import (
	"mime/multipart"
	"gorm.io/gorm"
)

type DbFile struct {
	gorm.Model
	Uuid string
	FileName string
}

type File struct {
	Content *multipart.FileHeader `form:"file" binding:"required"`
	Uuid string `form:"uuid"`
}

type Success struct {
	Uuid    string
	Message string
}
