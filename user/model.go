package user

type User struct {
	Id              int
	Name            string
	Age             int
	JobTitle        string
	Image           string
	Characteristics []string
	Likes           []string
	Dislikes        []string
	BackgroundColor string
	NameFontColor   string
	JobFontColor    string
	AgeFontColor    string
	About           string
}
