package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name            string
	Age             int
	JobTitle        string
	Image           string
	Characteristics []string `gorm:"-"`
	Likes           []string `gorm:"-"`
	Dislikes        []string `gorm:"-"`
	BackgroundColor string
	NameFontColor   string
	JobFontColor    string
	AgeFontColor    string
	About           string
}
