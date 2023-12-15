package db

import (
	"github.com/kubestaff/golearning/setting"
	"gorm.io/gorm"
)

func Migrate(dbConnection *gorm.DB) error  {
	dbConnection.AutoMigrate(&setting.UserSetting{})
	return nil
}

