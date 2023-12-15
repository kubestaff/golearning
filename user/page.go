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
