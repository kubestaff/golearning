package file

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kubestaff/golearning/apperror"
	"github.com/kubestaff/web-helper/server"
	"gorm.io/gorm"
)

type Handler struct {
	DbConnection *gorm.DB
}

func (h *Handler) Upload(c *gin.Context) {
	file := File{}

	if err := c.ShouldBind(&file); err != nil {
		c.JSON(400, server.JsonError{
			Error: err.Error(),
			Code:  400,
		})
		return
	}

	uuid, err := SaveFile(c, file, h.DbConnection, "avatar.png")
	if err != nil {
		if errors.Is(err, apperror.ErrInvalidInput) {
			c.JSON(400, server.JsonError{
				Error: err.Error(),
				Code:  400,
			})
			return
		}
		c.JSON(500, server.JsonError{
			Error: fmt.Sprintf("failed to save file %s", file.Content.Filename),
			Code:  500,
		})
		return
	}
	
	c.JSON(200, Success{
		Uuid: uuid,
		Message: "Successfully saved file " + file.Content.Filename,
	})
}