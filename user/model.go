package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name            string
	Age             int
	JobTitle        string
	Image           string
	Characteristics []string `gorm:"type:TEXT[]"`
	Likes           []string `gorm:"type:TEXT[]"`
	Dislikes        []string `gorm:"type:TEXT[]"`
	BackgroundColor string
	NameFontColor   string
	JobFontColor    string
	AgeFontColor    string
	About           string
}
