package user

type Provider struct {}

func (p Provider) GetAll() []User {
	return []User {
		{
			Id: 1,
			Name: "DR3",
			Age: 23,
			JobTitle: "pharmacist",
			Characteristics: []string{
				"Red hair",
				"Brown eyes",
			},
			Likes: []string{
				"F1",
				"Horses",
			},
			Dislikes: []string {
				"Mushroom",
				"Rude people",
			},
			BackgroundColour: "#3EBBC2",
			JobFontColour: "#4741BE",
			About: "I love watching movies and being with family.",
		},
		{
			Id: 2,
			Name: "Abrar",
			Age: 26,
			JobTitle: "Lawyer",
			Characteristics: []string{
				"Blonde hair",
				"Brown eyes",
			},
			Likes: []string{
				"Going on holidays",
				"Horses",
			},
			Dislikes: []string {
				"Mushroom",
				"Rude people",
			},
			BackgroundColour: "#3EBBC2",
			JobFontColour: "#4741BE",
			About: "I love watching movies and being with family.",
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
