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
			JobFontColor: "#EBE4F7",
			About: "I like shopping with Tiffany.",
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
		JobFontColor: "#B0709A",
		About: "I love to hike",
	},
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
