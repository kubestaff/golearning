package user

import (
	"os"
	"encoding/json"

	"github.com/kubestaff/golearning/helper"
)

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

func (p Provider) SaveUser(user *User, Id int) error {
	users, err := p.GetAll()
	if err!= nil {
		return err
	}

	for _, usr := range users{
	if usr.Id != Id{
		helper.SaveJSONFile(FileName, user)
		}
	
		jsonData, err := json.Marshal(user)
		if err != nil {
			return err
		}
	
		err = os.WriteFile(FileName, jsonData, 0644)
	if err != nil {
		return err
	}

	}

	return nil
}
	//find a user
	//if user is found replace it in the file
	//if user is not found add it at the bottom
	//if file cannot be saved, return an error