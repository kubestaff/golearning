package setting

import (
	"errors"
	"gorm.io/gorm"
)

type Provider struct {
	DbConnection *gorm.DB
}

func (p Provider) GetAllSettings() ([]UserSetting, error) {
	var settings []UserSetting
	result := p.DbConnection.Find(&settings)
	if result.Error != nil {
		return nil, result.Error
	}
	return settings, nil
}

func (p Provider) GetSettingByUserId(userId int) (s UserSetting, isFound bool, err error) {
	var setting UserSetting
	result := p.DbConnection.First(&setting, "user_id = ?", userId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return UserSetting{}, false, nil
	}

	if result.Error != nil {
		return UserSetting{}, false, nil
	}
	return setting, true, nil

}

func (p Provider) SaveSetting (newSetting *UserSetting) error {
	result := p.DbConnection.Save(newSetting)
	if result.Error != nil {
		return result.Error
	}
	return nil
}