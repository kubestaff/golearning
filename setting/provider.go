package setting

import (
	"errors"
	"os"

	"github.com/kubestaff/golearning/helpers"
	"github.com/kubestaff/golearning/setting"
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
	if newSetting.Id != 0 {
		return p.insertSetting(newSetting)
	}

	return p.updateSetting(newSetting)
}

func (p Provider) insertSetting(setting *UserSetting) error {
	existingSettings, err := p.GetAllSettings()
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

func (p Provider) updateSetting(setting *UserSetting) error {
	return nil
}
