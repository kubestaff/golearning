package user

type Provider struct {}

func (p Provider) GetAll() []User{
	return []User {
		{
			Id: 1,
			Name: "Max Verstappen",
			Age: 26,
			JobTitle: "F1 Driver",
			Characteristics: []string{
				"brown hair",
				"pink cheeks",
				"green eyes",
			},

			Likes: []string{
				"Driving",
				"Golf",
				"Boxing",
			},

			DisLikes: []string{
				"mushroom",
				"people who don't know how to drive",
			},

			BackgroundColour: "#996518",
			JobFrontColour: "#EDC78D",
			About: "RedBull didnt take me to victory I took RedBull to victory",
		},

		{
			Id: 2,
			Name: "Daniel Ricciardo",
			Age: 35,
			JobTitle: "F1 Driver",
			Characteristics: []string{
				"brown hair",
				"green eyes",
			},

			Likes: []string{
				"Driving",
				"Singing",
			},

			DisLikes: []string{
				"Pizza",
				"People",
			},

			BackgroundColour: "#996518",
			JobFrontColour: "#EDC78D",
			About: "I like music",
		},
	}
}

func(p Provider) GetUserById(id int)(user User, isFound bool){
	users:= p.GetAll()

	for _, user := range users {
		if user.Id == id{
			return user, true 
		}
	}
	return User{}, false
}
