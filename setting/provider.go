package setting

import (
	"errors"
	"os"

	"github.com/kubestaff/golearning/helpers"
)

type Provider struct{}

const fileName = "data/settings.json"

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
	if newSetting.ID == 0 {
		return p.insertSetting(newSetting)
	}
	return p.updateSetting(newSetting)
}

func (p Provider) insertSetting(setting *UserSetting) error {
	existingSettings, err := p.GetAll()
	if os.IsNotExist(err) {
		setting.ID = uint(len(existingSettings) + 1)
		settingsToSave := []UserSetting{
			*setting,
		}
		return p.SaveSettings(&settingsToSave)
	}
	if err != nil {
		return err
	}
	setting.ID = uint(len(existingSettings) + 1)

	existingSettings = append(existingSettings, *setting)
	return p.SaveSettings(&existingSettings)
}

// implement the updateSetting function
func (p Provider) updateSetting(setting *UserSetting) error {
	existingSettings, err := p.GetAll()
	if err != nil {
		return err
	}

	_, foundIndex, found := p.findSettingsById(existingSettings, int(setting.ID))
	if !found {
		return errors.New("setting not found to update")
	}

	foundSetting := existingSettings[foundIndex]
	foundSetting.UserId = setting.UserId

	if foundSetting.AmountOfUsersOnMainPage != setting.AmountOfUsersOnMainPage {
		foundSetting.AmountOfUsersOnMainPage = setting.AmountOfUsersOnMainPage
	}

	existingSettings[foundIndex] = foundSetting

	return p.SaveSettings(&existingSettings)

}

func (p Provider) findSettingsById(settings []UserSetting, id int) (*UserSetting, int, bool) {
	for i, setting := range settings {
		if int(setting.ID) == id {
			return &setting, i, true
		}
	}
	return nil, -1, false
}


//int -3, -2, -1, 0, 1, 2, 3
//uint 0