package user

import (
	"strconv"

	"github.com/kubestaff/golearning/helpers"
	"github.com/kubestaff/web-helper/server"
)


func HandleMe(inputs server.Input) (filename string, placeholders map[string]string) {
	//get id of user from url
	userIdStr := inputs.Values.Get("id")
	//convert id from string to int
	userIdInt, err := strconv.Atoi(userIdStr)
	if err != nil {
		return helpers.HandleErrorText("invalid user id")
	}

	//get user from id
	usersProvider := Provider{}
	user, isFound := usersProvider.GetUserById(userIdInt)
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
