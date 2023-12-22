package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

const Database = "golearning.db"

func CreateDatabase() (dbConnection *gorm.DB, err error) {
	dbConnection, err = gorm.Open(sqlite.Open(Database), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return dbConnection, err
}
