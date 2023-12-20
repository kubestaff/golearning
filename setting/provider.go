package setting

import (
	"errors"
	"os"

	"github.com/kubestaff/golearning/helpers"
)

type Provider struct{}

const fileName = "data/settings.json"

func (p Provider) GetAllSettings() ([]UserSetting, error) {
	var settings []UserSetting
	err := helpers.ReadFromJSONFile(fileName, &settings)
	if err != nil {
		return nil, err
	}
	return settings, nil
}

func (p Provider) GetSettingByUserId(userId int) (s UserSetting, isFound bool, err error) {
	settings, err := p.GetAllSettings()
	if err != nil {
		return UserSetting{}, false, err
	}
	for _, setting := range settings {
		if setting.Id == userId {
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
	existingSettings, err := p.GetAllSettings()
	if os.IsNotExist(err) {
		setting.Id = len(existingSettings) + 1
		settingsToSave := []UserSetting{
			*setting,
		}
		return p.SaveSettings(&settingsToSave)
	}
	if err != nil {
		return err
	}

	setting.Id = len(existingSettings) + 1
	existingSettings = append(existingSettings, *setting)
	return p.SaveSettings(&existingSettings)

}

func (p Provider) updateSetting(setting *UserSetting) error {
	existingSettings, err := p.GetAllSettings()
	if err != nil {
		return err
	}

	_, foundIndex, found := p.findSettingById(existingSettings, setting.Id)
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

func (p Provider) findSettingById(settings []UserSetting, id int) (*UserSetting, int, bool) {
	for i, setting := range settings {
		if setting.Id == id {
			return &setting, i, true
		}
	}
	return nil, -1, false
}
