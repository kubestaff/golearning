package main

import (
	"strconv"

	"github.com/kubestaff/golearning/helpers"
	"github.com/kubestaff/golearning/user"
	"github.com/kubestaff/web-helper/server"
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
	if err != nil {
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

	characteristicsStr := helpers.WrapStringsToTags(user.Characteristics, "li")
	likesStr := helpers.WrapStringsToTags(user.Likes, "li")
	dislikesStr := helpers.WrapStringsToTags(user.Dislikes, "li")

	output := map[string]string{
		"%name%":            user.Name,
		"%job-title%":       user.JobTitle,
		"%image%":           user.Image,
		"%age%":             strconv.Itoa(user.Age),
		"%about%":           user.About,
		"%characteristics%": characteristicsStr,
		"%likes%":           likesStr,
		"%dislikes%":        dislikesStr,
		"%about":            user.About,
		"%backgroundColor%": user.BackgroundColor,
		"%jobFontColor%":    user.JobFontColor,
		"%ageFontColor%":    user.AgeFontColor,
	}

	//output the user

	return "html/me.html", output
}

func HandleIndex(inputs server.Input) (filename string, placeholders map[string]string) {
	variables := map[string]string{"%name%": "Max Mustermann"}
	return "html/index.html", variables
}
