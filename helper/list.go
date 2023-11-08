package helper

import (
	"fmt"
	"strings"
)

func WrapStringsToTags(inputStrings []string, tag string) string {
	taggedInput := []string{}
	for _, inputString := range inputStrings {
		taggedInputString := fmt.Sprintf("<%s>%s</%s>", tag, inputString, tag)
		taggedInput = append(taggedInput, taggedInputString)
	}

	taggedInputsFlat := strings.Join(taggedInput, "")

	return taggedInputsFlat
}
