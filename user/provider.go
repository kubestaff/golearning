package user

import "github.com/kubestaff/golearning/helper"

type Provider struct{}

const FileName = "data/userData.json"

func (p Provider) GetAll() ([]User, error) {
	users := []User{}
	err := helper.ReadFromJSONFile(FileName, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (p Provider) GetUserById(id int) (usr User, isFound bool, err error) {
	users, err := p.GetAll()
	if err != nil {
		return User{}, false, err
	}

	for _, user := range users {
		if user.Id == id {
			return user, true, nil
		}
	}

	return User{}, false, nil
}

func (p Provider) SaveUsers() error {
	users := []User{
		{
			Id:       1,
			Name:     "Funto Awoyelu",
			Age:      27,
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
			About:           "This is a part of my golearning project",
			AgeFontColor:    "#000000",
		},
		{
			Id:              2,
			Name:            "Farah",
			Age:             26,
			JobTitle:        "Aspiring software engineer",
			Image:           "pexels-photo-992734.jpeg",
			BackgroundColor: "#000000",
			NameFontColor:   "#952424",
			JobFontColor:    "#ffffff",
			AgeFontColor:    "#ffffff",
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
			AgeFontColor:    "#9fd116",
			About:           "I have built this page as part of my golearning bootcamp.",
		},
		{
			Id:              4,
			Name:            "Sam Jobs",
			Age:             33,
			JobTitle:        "CEO",
			Image:           "spidermen.jpeg",
			BackgroundColor: "#333",
			NameFontColor:   "#fff",
			JobFontColor:    "#777",
			AgeFontColor:    "#9fd116",
			About:           "I'm Sam",
		},
		{
			Id:              5,
			Name:            "Matt Smith",
			Age:             38,
			JobTitle:        "Software Developer",
			Image:           "mattsmith.png",
			Characteristics: []string{
				"White hair",
				"Dark brown eyes",
			},
			Likes: []string{
				"Good food",
				"Long walks",
			},
			Dislikes: []string{
				"Cats",
				"Fireworks",
			},
			BackgroundColor: "#333",
			NameFontColor:   "#fff",
			JobFontColor:    "#777",
			AgeFontColor:    "#9fd116",
			About:           "I'm Matt Smith",
		},
	}

	return helper.SaveJSONFile(FileName, users)
}
