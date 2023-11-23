package user

<<<<<<< HEAD
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
=======
type Provider struct{}

func (p Provider) GetAll() []User {
	return []User {
		{
			Id: 1,
			Name: "Oreva",
			Age: 88,
			JobTitle: "programmer",
			Image: "headshot.jpeg",
			Characteristics: []string{
				"black hair",
				"brown eyes",
			},
			Likes: []string{
				"Poetry",
				"football",
			},
			Dislikes: []string{
				"dishonesty",
				"slime",
			},
			BackgroundColor: "#ffff",
			JobFontColor: "#777",
			About: "i love liverpool football club",
		},
		{

			Id: 2,
			Name: "Ansel",
			Age: 78,
			JobTitle: "Software dev",
			Image: "siren.png",
			Characteristics: []string{
				"white hair",
				"black eyes",
			},
			Likes: []string{
				"skiing",
				"football",
			},
			Dislikes: []string{
				"lies",
				"slime",
			},
			BackgroundColor: "#f6e3d4",
			JobFontColor: "#777",
			About: "i love to skate",
>>>>>>> oreva-main
		},
	}
}

<<<<<<< HEAD
func(p Provider) GetUserById(id int)(user User, isFound bool){
	users:= p.GetAll()

	for _, user := range users {
		if user.Id == id{
			return user, true 
		}
	}
	return User{}, false
}
=======
func (p Provider) GetUserById(id int) (usr User, isFound bool) {
	users := p.GetAll()

	for _, user := range users {
		if user.Id == id {
			return user, true
		}
	}
	return User{}, false
}
>>>>>>> oreva-main
