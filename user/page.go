package user

import (
	"github.com/kubestaff/web-helper/server"
	"strconv"
)

func HandleMe(inputs server.Input) (filename string, placeholders map[string]string) {
	userIdStr := inputs.Values.Get("id")
	userIdInt, err := strconv.Atoi(userIdStr)
	if err != nil {
		// todo
		return "", nil
	}

	usersProvider := Provider{}
	user, isFound := usersProvider.GetUserById(userIdInt)
	if !isFound {
		// todo
		return "", nil
	}

	output := map[string]string{
		"%name%":             user.Name,
		"%job-title%":        user.JobTitle,
		"%age%":              strconv.Itoa(user.Age),
		"%image%":            user.Image,
		"%characteristics0%": user.Characteristics[0],
		"%characteristics1%": user.Characteristics[1],
		"%characteristics2%": user.Characteristics[2],
		"%likes0%":           user.Likes[0],
		"%likes1%":           user.Likes[1],
		"%likes2%":           user.Likes[2],
		"%likes3%":           user.Likes[3],
		"%likes4%":           user.Likes[4],
		"%dislikes0%":        user.Dislikes[0],
		"%dislikes1%":        user.Dislikes[1],
		"%about%":            user.About,
	}

	return "html/me.html", output
}
