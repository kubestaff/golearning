package db

import (
	"github.com/kubestaff/golearning/user"
	"gorm.io/gorm"
)


func Migrate(dbConnection *gorm.DB) error {
	dbConnection.AutoMigrate(&user.User{})
	return nil
}