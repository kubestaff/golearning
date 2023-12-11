package file

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const UploadPath = "static/uploads"

func SaveFile(c *gin.Context, file File, dbConn *gorm.DB, fileName string) (string, error) {
	uniqueFolderId := ""
	dbFile := DbFile{}
	if file.Uuid != "" {
		dbConn.Find(&dbFile, "uuid = ?", file.Uuid)
		if dbFile.ID == 0 {
			return "", fmt.Errorf("invalid uuid: %s for file %s", file.Uuid, file.Content.Filename)
		}

		tempPath := filepath.Join(UploadPath, dbFile.Uuid)
		if _, err := os.Stat(tempPath); os.IsNotExist(err) {
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
