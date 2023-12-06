package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


const DataBaseName = "golearning.db"

func CreateDatabase() (dbConnection *gorm.DB, err error) {
	dbConnection, err = gorm.Open(sqlite.Open(DataBaseName), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return dbConnection, err
}