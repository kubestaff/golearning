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
	userIdStr := inputs.Values.Get("id")
	userIdInt, err := strconv.Atoi(userIdStr)
	if err != nil {
		return helpers.HandleErrorText("Invalid user id")
	}

	userProvider := Provider{
		DbConnection: h.DbConnection,
	}
	user, isFound, err := userProvider.GetUserById(userIdInt)
	if err != nil {
		return helpers.HandleErr(err)
	}
	if !isFound {
		return helpers.HandleErrorText("User not found")
	}

	characteristicsStr := helpers.WrapStringsToTags(user.Characteristics, "li")
	likesStr := helpers.WrapStringsToTags(user.Likes, "li")
	dislikesStr := helpers.WrapStringsToTags(user.Dislikes, "li")

	output := map[string]string{
		"%name%":            user.Name,
		"%job-title%":       user.JobTitle,
		"%age%":             strconv.Itoa(user.Age),
		"%image%":           user.Image,
		"%characteristics%": characteristicsStr,
		"%likes%":           likesStr,
		"%dislikes%":        dislikesStr,
		"%about%":           user.About,
		"%backgroundColor%": user.BackgroundColor,
		"%nameFontColor%:":  user.NameFontColor,
		"%jobFontColor%":    user.JobFontColor,
		"%ageFontColor%":    user.AgeFontColor,
	}
	return "html/me.html", output
}

func (h Handler) HandleReadUser(inputs server.Input) (filename string, placeholders map[string]string) {
	//if we have a name and age, we want to create a user
	if inputs.Get("name") != "" || inputs.Get("age") != "" {
		return h.HandleCreateOrUpdate(inputs)
	}

	output := map[string]string{
		"%id%":   "",
		"%name%": "",
		"%age%":  "",
	}
	//if we have an id, we just want to read the user from the database
	if inputs.Get("id") != "" {
		userIdStr := inputs.Values.Get("id")
		userIdInt, err := strconv.Atoi(userIdStr)
		if err != nil {
			return helpers.HandleErrorText("invalid user id")
		}
		userProvider := Provider{
			DbConnection: h.DbConnection,
		}
		//fetch the user by id
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

func (h Handler) HandleCreateOrUpdate(inputs server.Input) (filename string, placeholders map[string]string) {
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
		Age:  ageInt,
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
		return helpers.HandleErrorText("user not found")
	}

	err = userProvider.DeleteUser(&usr)
	if err != nil {
		return helpers.HandleErr(err)
	}

	output := map[string]string{
		"%success%": "User deleted successfully",
	}

	return "html/success.html", output
}
