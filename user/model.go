package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name            string
	Age             int
	JobTitle        string
	Image           string
	Characteristics []string `gorm:"type:text" db:"characteristics"`
	Likes           []string `gorm:"type:text" db:"likes"`
	Dislikes        []string `gorm:"type:text" db:"dislikes"`
	BackgroundColor string
	NameFontColor   string
	JobFontColor    string
	AgeFontColor    string
	About           string
}
