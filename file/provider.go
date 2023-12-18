package file

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kubestaff/golearning/apperror"
	"gorm.io/gorm"
)

const UploadPath = "static/uploads"
const UrlPath = "/static/uploads"

func SaveFile(c *gin.Context, file File, dbConn *gorm.DB, fileName string) (string, error) {
	uniqueFolderId := ""
	dbFile := DbFile{}
	if file.Uuid != "" {
		dbConn.Find(&dbFile, "uuid = ?", file.Uuid)
		if dbFile.ID == 0 {
			return "", fmt.Errorf("invalid uuid: %s for file %s: %w", file.Uuid, file.Content.Filename, apperror.ErrInvalidInput)
		}

		tempPath := filepath.Join(UploadPath, dbFile.Uuid)
		if _, e := os.Stat(tempPath); os.IsNotExist(e) {
			return "", fmt.Errorf("invalid uuid: %s for file %s", file.Uuid, file.Content.Filename)
		}

		dbFile.FileName = fileName
	} else {
		uniqueFolderId = uuid.NewString()
		dbFile.Uuid = uniqueFolderId
		dbFile.FileName = fileName
	}

	path := filepath.Join(UploadPath, dbFile.Uuid, fileName)

	err := c.SaveUploadedFile(file.Content, path)
	if err != nil {
		return "", err
	}

	result := dbConn.Save(&dbFile)

	err = result.Error
	if err != nil {
		os.Remove(path)
		return "", err
	}

	return dbFile.Uuid, err
}

func GetPathByUuid(uuid string, dbConn *gorm.DB) (string, error) {
	dbFile := DbFile{}
	dbConn.Find(&dbFile, "uuid = ?", uuid)
	if dbFile.ID == 0 {
		return "", nil
	}

	return filepath.Join(UploadPath, dbFile.Uuid, dbFile.FileName), nil
}

func GetUrlByUuid(domain, uuid string, dbConn *gorm.DB) (string, error) {
	dbFile := DbFile{}
	dbConn.Find(&dbFile, "uuid = ?", uuid)
	if dbFile.ID == 0 {
		return "", nil
	}

	return url.JoinPath(domain, UrlPath, dbFile.Uuid, dbFile.FileName)
}
