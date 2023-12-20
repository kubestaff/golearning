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

	for _, user := range users {
		if user.Id == id {
			return user, true, nil
		}
	}
	return User{}, false, nil
}

func (p Provider) SaveUsers(users *[]User) error {
	return helpers.SaveJSONFile(fileName, users)
}

func (p Provider) SaveUser(user User) error {
	users, err := p.GetAll()
	if err != nil {
		return err
	}

	for i, existingUser := range users {
		if existingUser.Id == user.Id {
			users[i] = user
			return helpers.SaveJSONFile(fileName, &users)
		}
	}
	users = append(users, user)
	return helpers.SaveJSONFile(fileName, &users)

}
