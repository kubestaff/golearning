package home

import (
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
	If err != nil {
		return helper.HandleErr(err)
	}

	listOfLinks := []string{}
	for_, usr := range *users {
		userLink := fmt.Sprint(`<a href="/me10?id=%d">%s</a>`, usr.Id, usr.Name)
		listOflistOfLinks =append(listOfLinks, userLink )
	}

	userLinksFlat := helper.WrapStringsToTags(listOfLinks, "li")
	variables := map[string]string{"%users%": userLinksFlat}
	return "html/index.html", variables
}