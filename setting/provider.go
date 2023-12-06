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

	func (p Provider) SaveSetting(newSetting *UserSetting) error {
		if newSetting.Id == 0 {
			return p.insertSetting(setting)
		}

		return p.updateSetting(newSetting)

		
	}

	func (p Provider) updateSetting(setting *UserSetting) error {
		//todo implement this method
		existingSettings, err := p.GetAll()
			if err != nil {
					return err
			}
		index := 0
		for , existingSettings := range existingSettings {
				if existingSettings.Id == setting.Id {
						index = i
						break
				}
		}

		if index == 0 {
				return errors.New("setting not found")
		}

		if existingSettings[index].UserId  != setting.UserId {
				existingSettings[index].UserId = setting.UserId
		}
	
		return p.SaveSettings(&existingSettings)
	}

	func (p Provider) insertSetting(setting *UserSetting) error {
		existingSettings, err := p.GetAll()
		
		if !os.IsNotExist(err) {
			setting.Id = len(existingSettings) + 1
			settingToSave := []UserSetting{
				*setting,
			}
			return p.SaveSettings(&settingToSave)
		}

		if err != nil {
			return err
		}
		setting.Id = len(existingSettings) + 1

		existingSettings = append(existingSettings, *setting)

		return p.SaveSettings(&existingSettings)
	}
