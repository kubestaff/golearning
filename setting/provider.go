package setting

import (
	"errors"
	"os"

	"github.com/kubestaff/golearning/helpers"
)

type Provider struct{}

const fileName = "data/userData.json"

func (p Provider) GetAll() ([]UserSetting, error) {
	var users []UserSetting
	err := helpers.ReadFromJSONFile(fileName, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (p Provider) GetSettingByUserId(userId int) (s UserSetting, isFound bool, err error) {
	settings, err := p.GetAll()
	if err != nil {
		return UserSetting{}, false, err
	}

	for _, setting := range settings {
		if setting.UserId == userId {
			return setting, true, nil
		}
	}
	return UserSetting{}, false, nil
}

func (p Provider) SaveSettings(settings *[]UserSetting) error {

	return helpers.SaveJSONFile(fileName, settings)
}

func (p Provider) SaveSetting(newSetting *UserSetting) error {
	if newSetting.Id == 0 {
		return p.insertSetting(newSetting)
	}
	return p.updateSetting(newSetting)
}

func (p Provider) insertSetting(setting *UserSetting) error {
	existingSettings, err := p.GetAll()
	if errors.Is(err, os.ErrNotExist) {
		settingsToSave := []UserSetting{
			*setting,
		}
		return p.SaveSettings(&settingsToSave)
	}
	if err != nil {
		return err
	}
	existingSettings = append(existingSettings, *setting)
	return p.SaveSettings(&existingSettings)
}

//implement the updateSetting function 
func (p Provider) updateSetting(setting *UserSetting) error {
	return nil
}
