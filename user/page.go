package user

import (
	"strconv"

	"github.com/kubestaff/golearning/helpers"
	"github.com/kubestaff/web-helper/server"
	"gorm.io/gorm"
)

type Handler struct {
	DbConnection *gorm.DB
}

func (h Handler) HandleMe(inputs server.Input) (filename string, placeholders map[string]string) {
	//get id of user from url
	userIdStr := inputs.Values.Get("id")
	//convert id from string to int
	userIdInt, err := strconv.Atoi(userIdStr)
	if err != nil {
		return helpers.HandleErrorText("invalid user id")
	}

	//get user from id
	usersProvider := Provider{
		DbConnection: h.DbConnection,
	}

	user, isFound, err := usersProvider.GetUserById(userIdInt)
	if err != nil {
		return helpers.HandleErr(err)
	}

	if !isFound {
		return helpers.HandleErrorText("user not found")
	}

	characteristicsStr := helpers.WrapStringsToTags(user.Characteristics, "li")
	likesStr := helpers.WrapStringsToTags(user.Likes, "li")
	dislikesStr := helpers.WrapStringsToTags(user.Dislikes, "li")

	output := map[string]string{
		"%name%":            user.Name,
		"%job-title%":       user.JobTitle,
		"%image%":           user.Image,
		"%age%":             strconv.Itoa(user.Age),
		"%characteristics%": characteristicsStr,
		"%likes%":           likesStr,
		"%dislikes%":        dislikesStr,
		"%backgroundColor%": user.BackgroundColor,
	}

	//output the user

	return "html/me.html", output
}

func (h Handler) HandleReadUser(inputs server.Input) (filename string, placeholders map[string]string) {
	if inputs.Get("name") != "" || inputs.Get("age") != "" {
		return h.HandleChangeUser(inputs)
	}

	output := map[string]string{
		"%id%":   "",
		"%name%": "",
		"%age%":  "",
	}

	if inputs.Get("id") != "" {
		userIdStr := inputs.Values.Get("id")
		userIdInt, err := strconv.Atoi(userIdStr)
		if err != nil {
			return helpers.HandleErrorText("invalid user id")
		}

		userProvider := Provider{
			DbConnection: h.DbConnection,
		}

		usr, isFound, err := userProvider.GetUserById(userIdInt)
		if err != nil {
			return helpers.HandleErr(err)
		}

		if isFound {
			output["%id%"] = userIdStr
			output["%name%"] = usr.Name
			output["%age%"] = strconv.Itoa(usr.Age)
		}

	}

	return "html/userform.html", output

}

func (h Handler) HandleChangeUser(inputs server.Input) (filename string, placeholders map[string]string) {
 	userIdStr := inputs.Values.Get("id")
	userIdInt := 0
	var err error

	if userIdStr != "" {
		userIdInt, err = strconv.Atoi(userIdStr)
		if err != nil {
			return helpers.HandleErrorText("invalid user id")
		}
	}

	name := inputs.Get("name")
	ageStr := inputs.Get("age")
	ageInt, err := strconv.Atoi(ageStr)
	if err != nil {
		return helpers.HandleErrorText("invalid age")
	}

	userProvider := Provider{
		DbConnection: h.DbConnection,
	}
	

	user := User{
		Name: name,
		Age: ageInt,
	}
	if userIdInt > 0 {
		user.ID = uint(userIdInt)
	}

	err = userProvider.SaveUser(&user)
	if err != nil {
		return helpers.HandleErr(err)
	}

	output := map[string]string{
		"%id%":   userIdStr,
		"%name%": name,
		"%age%":  ageStr,
	}

	return "html/userform.html", output

}

func (h Handler) HandleDeleteUser(inputs server.Input) (filename string, placeholders map[string]string) {
	userIdStr := inputs.Values.Get("id")
	userIdInt := 0
	var err error

	if userIdStr != "" {
		userIdInt, err = strconv.Atoi(userIdStr)
		if err != nil {
			return helpers.HandleErrorText("invalid user id")
		}
	}

	userProvider := Provider{
		DbConnection: h.DbConnection,
	}

	 usr, isFound, err := userProvider.GetUserById(userIdInt)
	 if err != nil {
		return helpers.HandleErr(err)
	 }

	 if !isFound {
		return helpers.HandleErrorText("user id is not found")
	 }

	 err = userProvider.DeleteUser(&usr)
	 if err != nil {
		return helpers.HandleErr(err)
	 }

	 output := map[string]string{
		"%success%": "Successfully deleted user",
	 }

	 return "html/success.html", output
	

}


