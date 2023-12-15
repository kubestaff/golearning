package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

const DatabaseSetting = "golearning.db"

func CreateDatatbase() (*gorm.DB, error) {
	dbConnection, err := gorm.Open(sqlite.Open("DatabaseSetting"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return dbConnection, err
}
