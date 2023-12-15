package setting

import "gorm.io/gorm"

type UserSetting struct {
	gorm.Model
	AmountOfUsersOnMainPage int
	UserId                  int
}
