package user

import (
	"strconv"
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
		//todo
		return "", nil
	}



	variables := map[string]string{
		"%name%": user.Name,
	}
	return "html/DR3.html", variables
}