package user

type Provider struct{}

func (p Provider) GetAll() []User {
	return []User{
		{
			Id:       1,
			Name:     "Funto Awoyelu",
			Age:      0,
			JobTitle: "Programme Manager",
			Image:    "funto.png",
			Characteristics: []string{
				"Dark brown hair",
				"Dark brown eyes",
				"5ft 3 height",
			},
			Likes: []string{
				"Reading",
				"Shopping",
			},
			Dislikes: []string{
				"Rudeness",
				"Celery",
			},
			BackgroundColor: "#f6e3d4",
			NameFontColor:   "#929522",
			JobFontColor:    "#777",
			About:           "",
		},
		{
			Id:              2,
			Name:            "Farah",
			Age:             26,
			JobTitle:        "Aspiring software engineer",
			Image:           "pexels-photo-992734.jpeg",
			BackgroundColor: "#000000",
			NameFontColor:   "#929524",
			JobFontColor:    "#778",
		},
		{
			Id:              3,
			Name:            "Moses Osho",
			Age:             26,
			JobTitle:        "Associate Developer",
			Image:           "spidermen.jpeg",
			BackgroundColor: "#333",
			NameFontColor:   "#fff",
			JobFontColor:    "#777",
			About:           "I have built this page as part of my golearning bootcamp.",
		},
	}
}

func (p Provider) GetUserById(id int) (usr User, isFound bool) {
	users := p.GetAll()
	for _, user := range users {
		if user.Id == id {
			return user, true
		}
	}

	return User{}, false
}