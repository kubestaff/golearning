package home

import (
	"fmt"

	"github.com/kubestaff/golearning/helper"
	"github.com/kubestaff/golearning/user"
	"github.com/kubestaff/web-helper/server"
)

//<a href="/me10?id=10">Yoni Makanda</a>
//convert
//<li><a href="/me10?id=10">Yoni Makanda</a></li>
func HandleHome(inputs server.Input) (filename string, placeholders map[string]string) {
	provider := user.Provider{}
	
	users, err:= provider.GetAll()
	if err != nil {
		return helper.HandleErr(err)
	}

	listOfLinks := []string{}
	for _, usr := range users {
		userLink := fmt.Sprintf(`<a href="/me10?id=%d">%s</a>`, int(usr.ID), usr.Name)
		listOfLinks =append(listOfLinks, userLink )
	}

	userLinksFlat := helper.WrapStringsToTags(listOfLinks, "li")
	variables := map[string]string{"%users%": userLinksFlat}
	return "html/index.html", variables
}