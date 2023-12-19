package user

import (
	"strconv"

	"github.com/kubestaff/golearning/helper"
	"github.com/kubestaff/web-helper/server"
	"gorm.io/gorm"
)

type Handler struct {
	DbConnection *gorm.DB
}

func (h Handler) HandleMe10(inputs server.Input) (server.Output) {
	userIdStr :=inputs.Values.Get("id")
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
	if inputs.Get("name") != ""  || inputs.Get("age") != "" || inputs.Get("job-title") !="" || inputs.Get("image") != "" ||  inputs.Get("about") !="" || inputs.Get("backgroundColor") !="" ||  inputs.Get("nameColor") !="" || inputs.Get("jobColor") !="" || inputs.Get("ageColor") !="" {
		return h.HandleChangeUser(inputs);
	}
	
	output := map[string]string{
		"%id%":				"",
		"%name%": 			"",
		"%age%": 			"0",
		"%job-title%": 		"",
		"%image%": 			"",
		"%characteristics%": "",
		"%likes%": 			"",
		"%dislikes%": 		"",
		"%about%": 			"",
		"%backgroundColor%": "",
		"%nameColor%": 		"",
		"%jobColor%": 		"",
		"%ageColor%": 		"",
	}

	if inputs.Get("id") != "" {
		userIdStr :=inputs.Values.Get("id")
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
		userIdStr :=inputs.Values.Get("id")

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
		jobtitle := inputs.Get("job-title")
		image := inputs.Get("image")
		characteristics := ("characteristics")
		likes := inputs.Get("likes")
		dislikes := inputs.Get("dislikes")
		about := inputs.Get("about")
		backgroundColor := inputs.Get("background_Color")
		namefontColor	:= inputs.Get("name_font_Color")	
		jobfontColor := inputs.Get("job_font_Color")
		agefontColor := inputs.Get("age_font_Color")	
		ageInt, err := strconv.Atoi(ageStr)
		if err != nil {
			return helper.HandleErrorText("Invalid age: a non-numeric value is provided: " + ageStr)
		}
	
		userProvider := Provider{
			DbConnection: h.DbConnection,
		}
		
		user := User{
			Name: name,
			Age: ageInt,
			JobTitle: jobtitle,
			Image: image,
			About: about,
			BackgroundColor: backgroundColor,
			NameFontColor: namefontColor,
			JobFontColor: jobfontColor,
			AgeFontColor: agefontColor,
			
		}

		if userIdInt > 0 {
			user.ID = uint(userIdInt)
		}

		err = userProvider.SaveUser(&user)
		if err != nil {
			return helper.HandleErr(err)
		}
	
		output := map[string]string{

			"%id%": 			 userIdStr,
			"%name%":			 name,
			"%age%":  			 ageStr,
			"%job-title%": 		 jobtitle,
			"%image%": 			 image,
			"%characteristics%": characteristics,
			"%about%": 			 about,
			"%likes%": 			 likes,
			"%dislikes%": 		 dislikes,
			"%backgroundColor%": backgroundColor,
			"%namefontColor%" :  namefontColor,
			"%jobfontColor%" :   jobfontColor,
			"%agefontColor%" : 	 agefontColor,
			
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
	usr, isFound, err:= userProvider.GetUserById(userIdInt)
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