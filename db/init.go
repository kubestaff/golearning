package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

const DatabaseName = "golearning.db"

func CreateDatabase() (*gorm.DB, error) {
	dbConnection, err := gorm.Open(sqlite.Open(DatabaseName), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return dbConnection, err
}