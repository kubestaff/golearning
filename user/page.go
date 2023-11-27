package user

import (
	"strconv"

	"github.com/kubestaff/golearning/helpers"
	"github.com/kubestaff/web-helper/server"
)

func HandleMe(inputs server.Input) (filename string, placeholders map[string]string) {
	userIdStr := inputs.Values.Get("id")
	userIdInt, err := strconv.Atoi(userIdStr)
	if err != nil {
		//todo
		return "", nil
	}

	userProvider := Provider{}
	user, isFound := userProvider.GetUserById(userIdInt)
	if !isFound {
		//todo
		return "", nil
	}

	characteristicsStr := helpers.WrapStringsToTags(user.Characteristics, "li")
	likesStr := helpers.WrapStringsToTags(user.Likes, "li")
	dislikesStr := helpers.WrapStringsToTags(user.Dislikes, "li")

	variables := map[string]string{
		"%name%":            user.Name,
		"%job-title%":       user.JobTitle,
		"%age%":             strconv.Itoa(user.Age),
		"%image%":           user.Image,
		"%characteristics%": characteristicsStr,
		"%likes%":           likesStr,
		"%dislikes%":        dislikesStr,
	}
	return "html/me.html", variables
}
