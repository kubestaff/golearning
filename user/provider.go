package user

import (


	"github.com/kubestaff/golearning/helper"
)

type Provider struct{}

const FileName = "data/userData.json"

func (p Provider) GetAll() (*[]User, error) {
	users := []User{}
	err := p.readUsersFromJsonFile(&users)
	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (p Provider) readUsersFromJsonFile(inputUsers *[]User) error {
	err := helper.ReadFromJSONFile(FileName, inputUsers)
	if err != nil {
		return err
	}

	return nil
}

func (p Provider) GetUserById(id int) (usr User, isFound bool, err error) {
	usersPointer, err := p.GetAll()
		if err!=nil {
			return user{}, false, err
		}

		for _, user := range *usersPointer {
			if int(user.ID) == id {
				return user, true, nil
		}
	}

	return User{}, false, nil

	
}
func (p Provider) SaveUsers() error {
	users = p.GetAll()
{
			Id:       10,
			Name:     "Yoni Makanda",
			Age:      37,
			JobTitle: "Aspiring Developer Contractor",
			Image:    "peofile.png",
			Characteristics: []string{
				"height:1m75",
				"eyes-color:brown",
				"hair-color:dark",
				"ethnicity:french-black",
			},
			Likes: []string{
				"Read",
				"Gym",
				"Music",
				"Coffee",
				"Good food",
			},
			Dislikes: []string{
				"To wait",
				"Corliflower",
			},
			BackgroundColor: "#faebd7",
			NameFontColor:   "#808080",
			JobFontColor:    "#777",
			About:           "",
			AgeFontColor:    "#000000",
		},
		{
			 Id: 		4,
			 Name: 		"Sam Jobs",
			 Age: 		33,
			 JobTitle: 	"CEO",
			 Image: 	"spidermen.jpeg",
			 Characteristics: null,
			 Likes: 	null,
			 Dislikes:  null,
			 BackgroundColor: "#333",
			 NameFontColor: "#fff",
			 JobFontColor: "#777",
			 AgeFontColor: "#9fd116",
			 About: "I'm Sam",
		},
		{
			 Id: 5,
			 Name: "Matt Smith",
			 Age: 38,
			 JobTitle: "Software Developer",
			 Image: "mattsmith.png",
			 Characteristics: []string{}
				 "White hair",
				 "Dark brown eyes",
		    },
		{	
			Likes: []string{}
				"Good food",
				"Long walks",
		    },
		{	
			Dislikes: []string{}
				"Cats",
				"Fireworks"
		   },
		{ 
			BackgroundColor: "#333",
			NameFontColor: "#fff",
			JobFontColor: "#777",
			AgeFontColor: "#9fd116",
			About: "I'm Matt Smith",
		  },
		
	


	

func (p Provider) SaveUsers(users *[]User) error {
	return helper.SaveJSONFile(FileName, users)
}

func (p Provider) SaveUsers(user *User) error {

	users, err := p.GetAll()

	if err != nil {

	}

	for i, existingUser := range users {
			if existingUser.Id == user.Id {
				users[i] = user
				return helper.SaveJSONFile(FileName, &users)
			}
	}

	users = append(users, user)
	return helper.SaveJSONFile(FileName, &users)
	//find a user
	//if user find replace it in the file
	//if user is not found add it at the bottom
	//if file cannot be saved, return error
}
