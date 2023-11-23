package user

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
