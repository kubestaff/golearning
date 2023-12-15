package user

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kubestaff/golearning/helper"
	"github.com/kubestaff/web-helper/server"
	"gorm.io/gorm"
)

// // add all user fields to the form
// read all users fields from the input
// save user with all input fields to the database
// check the result in the SQLite3 Editor VS Code extension

type Handler struct {
	DbConnection *gorm.DB
}

type UserExt struct {
	User
	ImagePath string
}

func (h Handler) HandleUser(c *gin.Context) {
	userIdStr := c.Query("id")
	userIdInt, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(400, server.JsonError{
			Error: fmt.Sprintf("Invalid number %s provided", userIdStr),
			Code:  400,
		})
		return
	}

	usersProvider := Provider{
		DbConnection: h.DbConnection,
	}
	user, isFound, err := usersProvider.GetUserById(userIdInt)
	if err != nil {
		c.JSON(500, server.JsonError{
			Error: err.Error(),
			Code:  500,
		})
		return
	}

	if !isFound {
		c.JSON(404, server.JsonError{
			Error: "User not found",
			Code:  404,
		})
		return
	}

	userEx := UserExt{
		User: user,
		ImagePath: "/lslsls",
	}

	c.JSON(200, userEx)
}

func (h Handler) HandleUsers(c *gin.Context) {
	userIdStr := c.Query("id")
	if userIdStr != "" {
		h.HandleUser(c)
		return
	}

	provider := Provider{
		DbConnection: h.DbConnection,
	}

	users, err := provider.GetAll()
	if err != nil {
		c.JSON(500, server.JsonError{
			Error: err.Error(),
			Code:  500,
		})
		return
	}
	c.JSON(200, users)
}

func (h Handler) HandleChangeUser(c *gin.Context) {
	userIdStr := c.Query("id")

	userIdInt := 0
	var err error
	if userIdStr != "" {
		userIdInt, err = strconv.Atoi(userIdStr)
		if err != nil {
			c.JSON(400, server.JsonError{
				Error: err.Error(),
				Code:  400,
			})
			return
		}
	}

	userFromInput := User{}
	if err := c.ShouldBindJSON(&userFromInput); err != nil {
		c.JSON(400, server.JsonError{
			Error: err.Error(),
			Code:  400,
		},
		)
		return
	}

	userProvider := Provider{
		DbConnection: h.DbConnection,
	}

	if userFromInput.Age < 0 || userFromInput.Age > 100 {
		c.JSON(400, server.JsonError{
			Error: fmt.Sprintf("Invalid age value %d", userFromInput.Age),
			Code:  400,
		})
		return
	}

	if userIdInt > 0 {
		userFromInput.ID = uint(userIdInt)
		userFromDB, isFound, err := userProvider.GetUserById(userIdInt)
		if err != nil {
			c.JSON(500, server.JsonError{
				Error: err.Error(),
				Code:  500,
			})
			return
		}

		if !isFound {
			c.JSON(404, server.JsonError{
				Error: "User id is not found",
				Code:  404,
			})
			return
		}

		if userFromInput.Name == "" {
			userFromInput.Name = userFromDB.Name
		}
		if userFromInput.About == "" {
			userFromInput.About = userFromDB.About
		}
		if userFromInput.Age == 0 {
			userFromInput.Age = userFromDB.Age
		}
		if userFromInput.JobTitle == "" {
			userFromInput.JobTitle = userFromDB.JobTitle
		}
	}

	err = userProvider.SaveUser(&userFromInput)
	if err != nil {
		c.JSON(500, server.JsonError{
			Error: err.Error(),
			Code:  500,
		})
		return
	}

	responseCode := 200
	message := "Successfully updated user"
	if userIdInt == 0 {
		responseCode = 201
		message = "Successfully created user"
	}

	c.JSON(responseCode, helper.Success{
		Message: message,
	})
}

func (h Handler) HandleDeleteUser(c *gin.Context) {
	userIdStr := c.Query("id")

	userIdInt := 0
	var err error
	if userIdStr != "" {
		userIdInt, err = strconv.Atoi(userIdStr)
		if err != nil {
			c.JSON(400, server.JsonError{
				Error: err.Error(),
				Code:  400,
			})
			return
		}
	}

	userProvider := Provider{
		DbConnection: h.DbConnection,
	}
	usr, isFound, err := userProvider.GetUserById(userIdInt)
	if err != nil {
		c.JSON(500, server.JsonError{
			Error: err.Error(),
			Code:  500,
		})
		return
	}

	if !isFound {
		c.JSON(404, server.JsonError{
			Error: "User id is not found",
			Code:  404,
		})
		return
	}

	err = userProvider.DeleteUser(&usr)
	if err != nil {
		c.JSON(500, server.JsonError{
			Error: err.Error(),
			Code:  500,
		})
		return
	}

	c.JSON(200, helper.Success{
		Message: fmt.Sprintf("Successfully deleted user %d", userIdInt),
	})
}
