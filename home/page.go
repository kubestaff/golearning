package home

import (
	"fmt"
<<<<<<< HEAD
=======

>>>>>>> oreva
	"github.com/kubestaff/golearning/helpers"
	"github.com/kubestaff/golearning/user"
	"github.com/kubestaff/web-helper/server"
)

<<<<<<< HEAD
// <li><a href="/me?id=1">Oreva</a>
// <li><a href="/me?id=2">Ansel</a>
// <li><a href="/me?id=3">Charles Leclerc</a>
// <li><a href="/me?id=4">Brittany Wilson</a>
// <li><a href="/me?id=5">Daniel Ricciardo</a>

func HandleHome(inputs server.Input) (filename string, placeholders map[string]string) {
	provider := user.Provider{}
	users := provider.GetAll()
	listOfLinks := []string{}
	for _, usr := range users {
		userlink := fmt.Sprintf(`<a href="/me?id=%d">%s</a>`, usr.Id, usr.Name)
		listOfLinks = append(listOfLinks, userlink)
	}
	userLinksFlat := helpers.WrapStringsToTags(listOfLinks, "Li")
	variables := map[string]string{"%users%": userLinksFlat}
=======
func HandleHome(inputs server.Input) (filename string, placeholders map[string]string) {
	provider := user.Provider{}

	users, err := provider.GetAll()
	if err != nil {
		return "html/error.html", nil
	}

	listOfLinks := []string{}

	for _, user := range users {
		userLink := fmt.Sprintf(`<a href="/me?id=%d">%s</a>`, user.Id, user.Name)
		listOfLinks = append(listOfLinks, userLink)
	}
	userLinkFlat := helpers.WrapStringsToTags(listOfLinks, "li")

	variables := map[string]string{"%users%": userLinkFlat}
>>>>>>> oreva
	return "html/index.html", variables
}
