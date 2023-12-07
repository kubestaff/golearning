package user

type Provider struct{}

func (p Provider) GetAll() []User {
	return []User{
		{
			Id: 1,
			Name: "Funto Awoyelu",
			Age: 26,
			JobTitle: "Programme Manager",
			Image: "funto.png",
			Characteristics: []string{
				"Dark brown hair",
				"Dark brown eyes",
				"5ft 3 height",
			},
			Likes: []string{
				"Shopping",
				"Good food",
				"Exploring in nature",
				"Meeting new people",
			},
			Dislikes: []string{
				"Rudeness",
				"Celery",
			},
			BackgroundColor: "#ccccff",
		},
		{
			Id: 2,
			Name: "Farah",
			Age: 26,
			JobTitle: "Aspiring software engineer",
			Image: "farah2.jpeg",
			About: "Creative marketeer looking to transition into tech and take up a new challenge.",
			Characteristics: []string{
				"Brown hair",
				"Brown eyes",
				"5ft 5 height",
			},
			Likes: []string{
				"Horse riding",
				"Reading",
				"Dancing",
				"Socializing",
			},
			Dislikes: []string{
				"Waking up early",
				"Brocolli",
			},
			BackgroundColor: "#99ccff",
		},
	}
}

func (p Provider) GetUserById(id int) (usr User, isFound bool) {
	users := p.GetAll()
	for _, user := range users{
		if user.Id == id {
			return user, true
		}
	}
	return User{}, false
}
