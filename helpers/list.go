package helpers

import (
	"fmt"
	"strings"
)

func WrapStringsToTags(inputStringsContainer []string, tag string) string {
	taggedInput := []string{}
	for _, individualInputString := range inputStringsContainer {
		taggedIndividualInputString := fmt.Sprintf("<%s>%s</%s>", tag, individualInputString, tag)
		taggedInput = append(taggedInput, taggedIndividualInputString)
	}
	joinedTaggedString := strings.Join(taggedInput, "")
	return joinedTaggedString
}
