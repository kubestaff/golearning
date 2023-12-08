package user

import "github.com/kubestaff/golearning/helpers"

type Provider struct{}

const fileName = "data/userData.json"

func (p Provider) GetAll() ([]User, error) {
	var users []User
	err := helpers.ReadFromJSONFile(fileName, &users)
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
	for _, user := range users{
		if user.Id == id {
			return user, true, nil
		}
	}
	return User{}, false, nil
}

func (p Provider) SaveUsers() error {
	users := []User{
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
			BackgroundColor: "#820D04",
		},
		{
			Id: 2,
			Name: "Farah",
			Age: 25,
			JobTitle: "Aspiring software engineer",
			Image: "farah.jpeg",
			Characteristics: []string{
				"brown hair",
				"brown eyes",
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
				"brocolli",
			},
			BackgroundColor: "#04825E",
		},
	}
	return helpers.SaveJSONFile(fileName, users)
}
