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

func (p Provider) SaveUsers(users *[]User) error {
	return helper.SaveJSONFile(FileName, users)
}

func (p Provider) SaveUser(user User) error {
	//find a user - get user by id
	users, err := p.GetAll()
	//if file cannot be saved, return an error
	if err != nil {
		return err
	}
	//if user is found replace it in the file - replace
	for i, existingUser := range users {
		if existingUser.Id == user.Id {
			users[i] = user
			return helper.SaveJSONFile(FileName, &users)
		}
	}
	//if user is not found add it at the bottom - append
	users = append(users, user)
	return helper.SaveJSONFile(FileName, &users)
}
