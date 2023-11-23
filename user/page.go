package user

import (
	"strconv"
<<<<<<< HEAD
    "github.com/kubestaff/web-helper/server"
)


func HandleMe(inputs server.Input) (filename string, placeholders map[string]string) {
	userIdStr:= inputs.Value.Get("id")
	userIdSInt, err:= strconv.Atoi(userIdStr)
	if err !=nil{
		//todo 
		return ""
	}


	userProvider:= Provider{}
	user, isFound := userProvider.GetUserById(userIdSInt)
	if !isFound{
=======

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
>>>>>>> oreva-main
		//todo
		return "", nil
	}

<<<<<<< HEAD


	variables := map[string]string{
		"%name%": user.Name,
	}
	return "html/DR3.html", variables
}
=======
	variables := map[string]string{
		"%name%":      user.Name,
		"%job-title%": user.JobTitle,
		"%age%":       strconv.Itoa(user.Age),
		"%image%":     user.Image,
	}
	return "html/me.html", variables
}
>>>>>>> oreva-main
