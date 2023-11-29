package setting

import (
	"errors"
	"gorm.io/gorm"
)

type Provider struct {
	DbConnection *gorm.DB
}

func (p Provider) GetAll() ([]UserSetting, error) {
	settings := []UserSetting{}
	result := p.DbConnection.Find(&settings)

	err := result.Error
	if err != nil {
		return nil, err
	}

	return settings, nil
}

func (p Provider) GetSettingByUserId(userId int) (s UserSetting, isFound bool, err error) {
	s = UserSetting{}

	result := p.DbConnection.First(&s, "user_id = ?", userId)

	err = result.Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return s, false, nil
	}

	if err != nil {
		return s, false, err
	}

	return s, true, nil
}

func (p Provider) SaveSetting(newSetting *UserSetting) error {
	result := p.DbConnection.Save(newSetting)

	err := result.Error
	if err != nil {
		return err
	}

	return nil
}

func (p Provider) DeleteSetting(newSetting *UserSetting) error { //Delete
	if newSetting.ID <= 0 {
		return errors.New("setting id not found")
	}
	result := p.DbConnection.Delete(newSetting, newSetting.ID)
	
	err := result.Error
	if err != nil {
		return err
	}

	return nil
}
