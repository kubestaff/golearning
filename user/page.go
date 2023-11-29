package user

import (
	
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

func (h Handler) HandleMe(inputs server.Input) server.Output {
	userIdStr := inputs.Values.Get("id")
	userIdInt, err := strconv.Atoi(userIdStr)
	if err != nil {
		return server.Output{
			Data: server.JsonError{
				Error: err.Error(),
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

func (h Handler) HandleReadUser(inputs server.Input) (filename string, placeholders map[string]string) {
	if inputs.Get("name") != "" || inputs.Get("age") != "" || inputs.Get("job-title") != "" || inputs.Get("image") != "" || inputs.Get("about") != "" || inputs.Get("backgroundColor") != "" || inputs.Get("nameColor") != "" || inputs.Get("jobColor") != "" || inputs.Get("ageColor") != "" {
		return h.HandleChangeUser(inputs)
	}

	output := map[string]string{
		"%id%":              "",
		"%name%":            "",
		"%age%":             "0",
		"%job-title%":       "",
		"%image%":           "",
		"%about%":           "",
		"%backgroundColor%": "",
		"%nameColor%":       "",
		"%jobColor%":        "",
		"%ageColor%":        "",
	}
	if inputs.Get("id") != "" {
		userIdStr := inputs.Values.Get("id")
		userIdInt, err := strconv.Atoi(userIdStr)
		if err != nil {
			return helper.HandleErrorText("Invalid user id")
		}

		userProvider := Provider{
			DbConnection: h.DbConnection,
		}

		usr, isFound, err := userProvider.GetUserById(userIdInt)
		if err != nil {
			return helper.HandleErr(err)
		}
		if isFound {
			output["%id%"] = userIdStr
			output["%name%"] = usr.Name
			output["%age%"] = strconv.Itoa(usr.Age)
			output["%job-title%"] = usr.JobTitle
			output["%image%"] = usr.Image
			output["%about%"] = usr.About
			output["%backgroundColor%"] = usr.BackgroundColor
			output["%nameColor%"] = usr.NameFontColor
			output["%jobColor%"] = usr.JobFontColor
			output["%ageColor%"] = usr.AgeFontColor
		}
	}

	return "html/userForm.html", output
}

func (h Handler) HandleChangeUser(inputs server.Input) (filename string, placeholders map[string]string) {
	userIdStr := inputs.Values.Get("id")

	userIdInt := 0
	var err error
	if userIdStr != "" {
		userIdInt, err = strconv.Atoi(userIdStr)
		if err != nil {
			return helper.HandleErrorText("Invalid user id")
		}
	}

	name := inputs.Get("name")
	ageStr := inputs.Get("age")
	jobTitle := inputs.Get("job-title")
	image := inputs.Get("image")
	about := inputs.Get("about")
	backgroundColor := inputs.Get("backgroundColor")
	nameColor := inputs.Get("nameColor")
	jobColor := inputs.Get("jobColor")
	ageColor := inputs.Get("ageColor")
	ageInt, err := strconv.Atoi(ageStr)
	if err != nil {
		return helper.HandleErrorText("Invalid age: a non-numeric value is provided: " + ageStr)
	}

	userProvider := Provider{
		DbConnection: h.DbConnection,
	}

	user := User{
		Name:            name,
		Age:             ageInt,
		JobTitle:        jobTitle,
		Image:           image,
		About:           about,
		BackgroundColor: backgroundColor,
		NameFontColor:   nameColor,
		JobFontColor:    jobColor,
		AgeFontColor:    ageColor,
	}

	if userIdInt > 0 {
		user.ID = uint(userIdInt)
	}

	err = userProvider.SaveUser(&user)
	if err != nil {
		return helper.HandleErr(err)
	}

	output := map[string]string{
		"%id%":              userIdStr,
		"%name%":            name,
		"%age%":             ageStr,
		"%job-title%":       jobTitle,
		"%image%":           image,
		"%about%":           about,
		"%backgroundColor%": backgroundColor,
		"%nameColor%":       nameColor,
		"%jobColor%":        jobColor,
		"%ageColor%":        ageColor,
	}

	return "html/userForm.html", output
}

func (h Handler) HandleDeleteUser(inputs server.Input) (filename string, placeholders map[string]string) {
	userIdStr := inputs.Values.Get("id")

	userIdInt := 0
	var err error
	if userIdStr != "" {
		userIdInt, err = strconv.Atoi(userIdStr)
		if err != nil {
			return helper.HandleErrorText("Invalid user id")
		}
	}

	userProvider := Provider{
		DbConnection: h.DbConnection,
	}
	usr, isFound, err := userProvider.GetUserById(userIdInt)
	if err != nil {
		return helper.HandleErr(err)
	}

	if !isFound {
		return helper.HandleErrorText("User id is not found")
	}

	err = userProvider.DeleteUser(&usr)
	if err != nil {
		return helper.HandleErr(err)
	}

	return "html/success.html", map[string]string{"%success%": "Successfully deleted user"}
}
