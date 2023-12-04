package home

import (
	"fmt"
	"github.com/kubestaff/golearning/helpers"
	"github.com/kubestaff/golearning/user"
	"github.com/kubestaff/web-helper/server"
)

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
	return "html/index.html", variables
}
