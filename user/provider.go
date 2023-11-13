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

func (p Provider) SaveUser(user *User) error {
	//find a user 
	//if user is found replace it in the file
	//if user is not found add it at the bottom
	//if file cannot be saved, return an error

	 users := []User{}

  
    // Check if the user already exists
    userExists := false
    for i, u := range users {
        if u.Id == user.Id { // User struct has the Id field in model.go file 
            users[i] = *user
            userExists = true
            break
        }
    }

    // If user doesn't exist, add to the list
    if !userExists {
        users = append(users, *user)
    }

    // Save the updated list back to the file
    return helper.SaveJSONFile(FileName, &users)
	
}