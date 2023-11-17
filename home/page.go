package home

import (
	"fmt"

	"github.com/kubestaff/golearning/helper"
	"github.com/kubestaff/golearning/user"
	"github.com/kubestaff/web-helper/server"
)

// <a href="/me?id=1">Funto Awoyelu</a>
// <a href="/me?id=2">Farah</a>
// <a href="/me?id=3">Moses Osho</a>
// convert
// <li><a href="/me?id=1">Funto Awoyelu</a></li><li><a href="/me?id=2">Farah</a></li><li><a href="/me?id=3">Moses Osho</a></li>
func HandleHome(inputs server.Input) (filename string, placeholders map[string]string) {
	provider := user.Provider{}

	users, err := provider.GetAll()
	if err != nil {
		return helper.HandleErr(err)
	}

	listOfLinks := []string{}
	for _, usr := range users {
		userLink := fmt.Sprintf(`<a href="/me?id=%d">%s</a>`, usr.ID, usr.Name)
		listOfLinks = append(listOfLinks, userLink)
	}

	userLinksFlat := helper.WrapStringsToTags(listOfLinks, "li")
	variables := map[string]string{"%users%": userLinksFlat}
	return "html/index.html", variables
}
