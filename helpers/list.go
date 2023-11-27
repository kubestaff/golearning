package helpers

import (
	"fmt"
	"strings"
)

func WrapStringsToTags(inputStrings []string, tag string) string {
	taggedInput := []string{}
	for _, inputString := range inputStrings {
		taggedInputString := fmt.Sprintf("<%s> %s </%s>",tag, inputString, tag )
		taggedInput = append(taggedInput, taggedInputString)
	}
	taggedInputFlat := strings.Join(taggedInput, "")
	return taggedInputFlat
}

// <li>Dark brown hair</li><li>Dark brown eyes</li><li>5ft 3 height </li>