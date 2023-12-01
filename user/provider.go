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
			BackgroundColor: "#271F30",
			NameFontColor: "#0000",
			JobFontColor: "#777",
			About: "i love liverpool football club",
			AgeFontColor: "#6C5A49",
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
			NameFontColor: "#110B06",
			JobFontColor: "#777",
			About: "i love to skate",
			AgeFontColor: "#F87F0E",
		},
		{

			Id: 3,
			Name: "Charles Leclerc",
			Age: 20,
			JobTitle: "F1 Ferarri Racer",
			Image: "cl16.png",
			Characteristics: []string{
				"brown hair",
				"green eyes",
			},
			Likes: []string{
				"F1",
				"Yachts",
			},
			Dislikes: []string{
				"Bad team strategist",
				"Losing",
			},
			BackgroundColor: "#E598A7",
			NameFontColor: "#070300",
			JobFontColor: "#8A2337",
			About: "I love F1",
			AgeFontColor: "#DAD0B1",
		},
		{

			Id: 4,
			Name: "Brittany Wilson",
			Age: 24,
			JobTitle: "Tycoon daughter",
			Image: "Wilson Sisters.png",
			Characteristics: []string{
				"blonde hair",
				"blue eyes",
			},
			Likes: []string{
				"Shopping",
				"Latrell Spencer",
			},
			Dislikes: []string{
				"Heather and Megan Vandergeld",
				"When baby poops",
			},
			BackgroundColor: "#6C3BC4",
			NameFontColor: "D0FCB3",
			JobFontColor: "#EBE4F7",
			About: "I like shopping with Tiffany.",
			AgeFontColor: "#609040",
		},
    {
		Id: 5,
		Name: "Daniel Ricciardo",
		Age: 34,
		JobTitle: "F1 Alpha Tauri driver",
		Image: "DR3.png",
		Characteristics: []string{
			"brown hair",
			"brown eyes",
		},
		Likes: []string{
			"Swimming",
			"Golf",
		},
		Dislikes: []string{
			"Bad drivers",
			"mushrooms",
		},
		BackgroundColor: "#CF7F9A",
		NameFontColor: "#58D505",
		JobFontColor: "#B0709A",
		About: "I love to hike",
		AgeFontColor: "#9BC59D",
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
