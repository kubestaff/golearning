package user

import "github.com/kubestaff/golearning/helper"

type Provider struct{
	DbConnection *gorm.DB
}

func (p Provider) GetAll() ([]User, error) {
	users := []User{}
	result := p.DbConnection.Find(&users)

	err := result.Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (p Provider) GetUserById(id int) (usr User, isFound bool, err error) {
	usr = User{}

	result := p.DbConnection.First(&usr, id)
	
	err = result.Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return usr, false, nil
	}

	if err != nil {
		return usr, false, err
	}

	return usr, true, nil
}

func (p Provider) SaveUser(usr *User) error {
	result := p.DbConnection.Save(usr)
	
	err := result.Error
	if err != nil {
		return err
	}

func (p Provider) SaveUser(user *User) error {
	//find a user
	//if user is found replace it in the file
	//if user is not found add it at the bottom
	//if file cannot be saved, return an error
	return nil
}