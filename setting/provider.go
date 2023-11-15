package setting

import (
	"errors"
	"os"

	"github.com/kubestaff/golearning/helper"
)

type Provider struct{}

const FileName = "data/settings.json"

func (p Provider) GetAll() ([]UserSetting, error) {
	users := []UserSetting{}
	err := helper.ReadFromJSONFile(FileName, &users)
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
	return helper.SaveJSONFile(FileName, settings)
}

func (p Provider) SaveSetting(newSetting *UserSetting) error {
	// 	find if this object already exists in the database (by finding a database object by id) or if id is 0 then go to insert mode
	// if you found an object in the database -> use update mode
	// if you didn't find anything with this id -> use insert mode
	// insert mode:
	// generate a new id if needed (check the amount of users in the database/slice and +1) len(slice)+1
	// append new data to the database (add at the bottom of the file/collection of objects etc)
	// update mode:
	// compare your existing object with your new object if there no changes just exist the update function
	// if there are changes then you should replace existing object with the new one
	if newSetting.Id != 0 {
		return p.insertSetting(newSetting)
	}

	return p.updateSetting(newSetting)
}

func (p Provider) updateSetting(setting *UserSetting) error {
	//todo implement this method
	return nil
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
