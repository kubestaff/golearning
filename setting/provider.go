package setting

import "github.com/kubestaff/golearning/helper"

type Provider struct{}

const FileName = "data/settings.json"

func (p Provider) GetAll() (UserSetting, error) {
	users := []UserSetting{}
	err := p.readUsersFromJsonFile(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (p Provider) GetSettingByUserId(userId int) (s UserSetting, isFound bool, err error) {
	settings, err := p.GetAll()
		if err!=nil {
			return UserSetting{}, false, err
		}
	}

		for _, setting := range users {
			if setting.userId == userId {
				return settng, true, nil
		}
	}

	return UserSetting{}, false, nil

	func (p Provider) SaveSettings(newSetting *UserSetting) error { 
		return helper.SaveJSONFile(FileName, settings) 

	}

	func (p Provider) SaveSetting(setting *UserSetting) error {
		if newSetting.Id != 0 {
			return p.insertSetting(setting)
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
			settingToSave := []UserSetting{
				*setting,
			}
			return p.SaveSettings(&settingToSave)
		}

		if err != nil {
			return err
		}

		existingSettings = append(existingSettings, *setting)

		return p.SaveSettings(&settingToSave)
	}
