package user

import (
	"strconv"

	"github.com/kubestaff/web-helper/server"
)

func HandleMe(inputs server.Input) (filename string, placeholders map[string]string) {
	userIdStr :=inputs.Values.Get("id")
	userIdInt, err := strconv.Atoi(userIdStr)
	if err !=nil {
		//todo
		return "", nil
	}

	userProvider := Provider{}
	user, isFound := userProvider.GetUserById(userIdInt)
	if !isFound {
		//todo
		return "", nil
	}

	variables := map[string]string{
		"%name%": user.Name,
	}
	return "html/me.html", variables
}