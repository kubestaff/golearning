package user

import (
	"image"
	"strconv"

	"github.com/kubestaff/golearning/helper"
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
		return helper.HandleErrorText("Invalid user id")
	}

	usersProvider := Provider{
		DbConnection: h.DbConnection,
	}
	user, isFound, err := usersProvider.GetUserById(userIdInt)
	if err != nil {
		return helper.HandleErr(err)
	}
	if !isFound {
		return helper.HandleErrorText("Not found")
	}

	//<li>characteristic[0]</li>
	//<li>characteristic[1]</li>
	//<li>characteristic[2]</li>
	//convert to:
	//<li>characteristic[0]</li><li>characteristic[1]</li><li>characteristic[2]</li>
	characteristicsStr := helper.WrapStringsToTags(user.Characteristics, "li")
	likesStr := helper.WrapStringsToTags(user.Likes, "li")
	disLikesStr := helper.WrapStringsToTags(user.Dislikes, "li")

	output := map[string]string{
		"%name%":            user.Name,
		"%job-title%":       user.JobTitle,
		"%age%":             strconv.Itoa(user.Age),
		"%image%":           user.Image,
		"%characteristics%": characteristicsStr,
		"%likes%":           likesStr,
		"%dislikes%":        disLikesStr,
		"%about%":           user.About,
		"%backgroundColor%": user.BackgroundColor,
		"%nameColor%":       user.NameFontColor,
		"%jobColor%":        user.JobFontColor,
		"%ageColor%":        user.AgeFontColor,
	}

	return "html/me.html", output
}

func (h Handler) HandleReadUser(inputs server.Input) (filename string, placeholders map[string]string) {
	
	if inputs.Get("name") != "" || inputs.Get("age") != "" {
		return h.HandleCreateUser(inputs)
	}

	output := map[string]string{
		"%name%": "",
		"%age%":  "0",
	}

	return "html/userForm.html", output
}

func (h Handler) HandleCreateUser(inputs server.Input) (filename string, placeholders map[string]string) {
	name := inputs.Get("name")
	ageStr := inputs.Get("age")
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
		JobTitle: JobTitle,
		Image: image,
		BackgroundColor: backgroundColor,
		NameFontColor: NameColor,
		JobFontColor: JobColor,
		AgeFontColor: AgeColor, 
		About: about,
	}
	err = userProvider.SaveUser(&user)

	if err != nil {
		return helper.HandleErr(err)
	}

	output := map[string]string{
		"%name%": name,
		"%age%":  ageStr,
	}

	return "html/userForm.html", output
}