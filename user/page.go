package user

import (
	"fmt"
	"strconv"

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

func (h Handler) HandleUser(inputs server.Input) server.Output {
	userIdStr := inputs.Values.Get("id")
	userIdInt, err := strconv.Atoi(userIdStr)
	if err != nil {
		return server.Output{
			Data: server.JsonError{
				Error: fmt.Sprintf("Invalid number %s provided", userIdStr),
				Code:  400,
			},
			Code: 400,
		}
	}

	usersProvider := Provider{
		DbConnection: h.DbConnection,
	}
	user, isFound, err := usersProvider.GetUserById(userIdInt)
	if err != nil {
		return server.Output{
			Data: server.JsonError{
				Error: err.Error(),
				Code:  500,
			},
			Code: 500,
		}
	}

	if !isFound {
		return server.Output{
			Data: server.JsonError{
				Error: "User not found",
				Code:  404,
			},
			Code: 404,
		}
	}

	return server.Output{
		Data: user,
		Code: 200,
	}
}

func (h Handler) HandleUsers(inputs server.Input) (o server.Output) {
	userIdStr := inputs.Values.Get("id")
	if userIdStr != "" {
		return h.HandleUser(inputs)
	}

	provider := Provider{
		DbConnection: h.DbConnection,
	}

	users, err := provider.GetAll()
	if err != nil {
		return server.Output{
			Data: server.JsonError{
				Error: err.Error(),
				Code:  500,
			},
			Code: 500,
		}
	}
	return server.Output{
		Data: users,
		Code: 200,
	}
}

func (h Handler) HandleChangeUser(serverInput server.Input) (output server.Output) {
	userIdStr := serverInput.Values.Get("id")

	userIdInt := 0
	var err error
	if userIdStr != "" {
		userIdInt, err = strconv.Atoi(userIdStr)
		if err != nil {
			return server.Output{
				Data: server.JsonError{
					Error: err.Error(),
					Code:  400,
				},
				Code: 400,
			}
		}
	}

	userFromInput := User{}
	err = serverInput.Scan(&userFromInput)
	if err != nil {
		return server.Output{
			Data: server.JsonError{
				Error: err.Error(),
				Code:  400,
			},
			Code: 400,
		}
	}

	userProvider := Provider{
		DbConnection: h.DbConnection,
	}

	if userFromInput.Age < 0 || userFromInput.Age > 100 {
		return server.Output{
			Data: server.JsonError{
				Error: fmt.Sprintf("Invalid age value %d", userFromInput.Age),
				Code:  400,
			},
			Code: 400,
		}
	}
	if userIdInt > 0 {
		userFromInput.ID = uint(userIdInt)
		userFromDB, isFound, err := userProvider.GetUserById(userIdInt)
		if err != nil {
			return server.Output{
				Data: server.JsonError{
					Error: err.Error(),
					Code:  500,
				},
				Code: 500,
			}
		}
		if !isFound {
			return server.Output{
				Data: server.JsonError{
					Error: "User id is not found",
					Code:  404,
				},
				Code: 404,
			}
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
		return server.Output{
			Data: server.JsonError{
				Error: err.Error(),
				Code:  500,
			},
			Code: 500,
		}
	}

	responseCode := 200
	message := "Successfully updated user"
	if userIdInt == 0 {
		responseCode = 201
		message = "Successfully created user"
	}

	return server.Output{
		Data: helper.Success{
			Message: message,
		},
		Code: responseCode,
	}
}

func (h Handler) HandleDeleteUser(inputs server.Input) (output server.Output) {
	userIdStr := inputs.Values.Get("id")

	userIdInt := 0
	var err error
	if userIdStr != "" {
		userIdInt, err = strconv.Atoi(userIdStr)
		if err != nil {
			return server.Output{
				Data: server.JsonError{
					Error: err.Error(),
					Code:  400,
				},
				Code: 400,
			}
		}
	}

	userProvider := Provider{
		DbConnection: h.DbConnection,
	}
	usr, isFound, err := userProvider.GetUserById(userIdInt)
	if err != nil {
		return server.Output{
			Data: server.JsonError{
				Error: err.Error(),
				Code:  500,
			},
			Code: 500,
		}
	}

	if !isFound {
		return server.Output{
			Data: server.JsonError{
				Error: "User id is not found",
				Code:  404,
			},
			Code: 404,
		}
	}

	err = userProvider.DeleteUser(&usr)
	if err != nil {
		return server.Output{
			Data: server.JsonError{
				Error: err.Error(),
				Code:  500,
			},
			Code: 500,
		}
	}

	return server.Output{
		Data: helper.Success{
			Message: fmt.Sprintf("Successfully deleted user %d", userIdInt),
		},
		Code: 200,
	}
}
