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

// func (p Provider) readUsersFromJsonFile(inputUsers *[]User) error {
// 	err := helper.ReadFromJSONFile(FileName, inputUsers)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func (p Provider) GetUserById(id int) (usr User, isFound bool, err error) {
	users, err := p.GetAll()
	if err != nil {
		return User{}, false, err
	}

	for _, user := range users {
		if int(user.ID) == id {
			return user, true, nil
		}
	}

	return User{}, false, nil
}

func (p Provider) SaveUsers(users *[]User) error {
	return helper.SaveJSONFile(FileName, users)
}

func (p Provider) SaveUser(user User) error {
	users, err := p.GetAll()
	if err != nil {
		return err
	}

	for i, existingUser := range users {
		if int(existingUser.ID) == int(user.ID) {
			users[i] = user
			return helper.SaveJSONFile(FileName, &users)
		}
	}

	users = append(users, user)
	return helper.SaveJSONFile(FileName, &users)

}
