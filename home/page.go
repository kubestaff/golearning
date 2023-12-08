package home

import (
	"fmt"

	"github.com/kubestaff/golearning/helpers"
	"github.com/kubestaff/golearning/user"
	"github.com/kubestaff/web-helper/server"
)

func HandleIndex(inputs server.Input) (filename string, placeholders map[string]string) {
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

	userLinksFlat := helpers.WrapStringsToTags(listOfLinks, "li")

	variables := map[string]string{"%user%": userLinksFlat}
	return "html/index.html", variables
}

// listOfLinks := {`<a href="/me?id=1">Funto</a>`, `<a href="/me?id=2">Farah</a>`}

// userLinksFlat := `<li><a href="/me?id=1">Funto</a></li><li><a href="/me?id=2">Farah</a></li>`
