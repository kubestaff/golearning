package user

import "github.com/kubestaff/web-helper/server"
import "strconv"

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
		"%name%": user.Name,
		"%job-title%": user.JobTitle,
	}

	return "html/me.html", output
}