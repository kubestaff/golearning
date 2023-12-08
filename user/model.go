package user

import (
	"gorm.io/gorm"
	"github.com/lib/pq"
)

type User struct {
	gorm.Model
	Name            string
	Age             int
	JobTitle        string
	Image           string
	Characteristics pq.StringArray `gorm:"type:TEXT[]"`
	Likes           pq.StringArray `gorm:"type:TEXT[]"`
	Dislikes        pq.StringArray `gorm:"type:TEXT[]"`
	BackgroundColor string
	NameFontColor   string
	JobFontColor    string
	AgeFontColor    string
	About           string
}
