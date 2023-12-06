package main

import (
	"strconv"

	"github.com/kubestaff/web-helper/server"
	"github.com/kubestaff/golearning/user"
)

func main() { 
	opts := server.Options{
		Port: 34567,
	}
	// we create the simplified web server
	s := server.NewServer(opts)

	// we close the server at the end
	defer s.Stop()

	s.Handle("/", HandleIndex)
	s.Handle("/me", HandleMe)


	s.Start()
}

func HandleMe(inputs server.Input) (filename string, placeholders map[string]string) {
	//get id of user from url
	userIdStr := inputs.Values.Get("id")
	//convert id from string to int
	userIdInt, err := strconv.Atoi(userIdStr)
	if err != nil{
		//todo
		return "", nil
	}
	
	//get user from id
	usersProvider := user.Provider{}
	user, isFound := usersProvider.GetUserById(userIdInt)
	if !isFound {
		//todo
		return "", nil
	}

	output := map[string]string{
		"%name%": user.Name,
		"%job-title%": user.JobTitle,
		"%age%": strconv.Itoa(user.Age),
		"%background-color%": user.BackgroundColor,
		"%job-font-color%": user.JobFontColor,
		"%age-font-color%": user.AgeFontColor,
		"%about": user.About,
		"%image%": user.Image,
	}
	
	//output the user

	return "html/me.html", output
}

func HandleIndex(inputs server.Input) (filename string, placeholders map[string]string) {
	variables := map[string]string{"%name%": "Max Mustermann"}
	return "html/index.html", variables
}
