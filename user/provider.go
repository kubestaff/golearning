package user

import (
	"errors"

	"gorm.io/gorm"
)

type Provider struct {
	DbConnection *gorm.DB
}

func (p Provider) GetAll() ([]User, error) {
	var users []User
	result := p.DbConnection.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (p Provider) GetUserById(id int) (usr User, isFound bool, err error) {
	var user User
	result := p.DbConnection.First(&user, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return User{}, false, nil
	}
	if result.Error != nil {
		return User{}, false, result.Error
	}

	return user, true, nil
}

func (p Provider) SaveUser(newUser *User) error {
	result := p.DbConnection.Save(newUser)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func (p Provider) DeleteUser(user *User) error {
	result := p.DbConnection.Delete(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}