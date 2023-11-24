package user

import (
	"errors"
	"gorm.io/gorm"
)

type Provider struct{
	DbConnection *gorm.DB
}

func (p Provider) GetAll() ([]User, error) { //Read
	users := []User{}
	result := p.DbConnection.Find(&users)

	err := result.Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (p Provider) GetUserById(id int) (usr User, isFound bool, err error) { //Read
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

func (p Provider) SaveUser(usr *User) error { //Update/Create
	result := p.DbConnection.Save(usr)
	
	err := result.Error
	if err != nil {
		return err
	}

	return nil
}

func (p Provider) DeleteUser(usr *User) error { //Delete
	result := p.DbConnection.Delete(usr, usr.ID)
	
	err := result.Error
	if err != nil {
		return err
	}

	return nil
}

