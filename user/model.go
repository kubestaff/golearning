package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name            string
	Age             int
	JobTitle        string
	Image           string
	Characteristics []string `gorm:"type:text"`
	Likes           []string `gorm:"type:text"`
	Dislikes        []string `gorm:"type:text"`
	BackgroundColor string
	NameFontColor   string
	JobFontColor    string
	AgeFontColor    string
	About           string
}
