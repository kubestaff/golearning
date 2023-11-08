package user

import (
	"strconv"

	"github.com/kubestaff/golearning/helper"
	"github.com/kubestaff/web-helper/server"
)

func HandleMe(inputs server.Input) (filename string, placeholders map[string]string) {
	userIdStr := inputs.Values.Get("id")
	userIdInt, err := strconv.Atoi(userIdStr)
	if err != nil {
		return helper.HandleErrorText("Invalid user id")
	}

	usersProvider := Provider{}
	user, isFound := usersProvider.GetUserById(userIdInt)
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
