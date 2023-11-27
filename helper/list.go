package helper

import (
	"fmt"
	"strings"
)

func WrapStringsToTags(inputStrings []string, tag string) string {
	taggedInput := []string{}
	for_, inputString := range inputinputStrings {
		taggedCharacteristic := fmt.Sprintf("<%s>%s</%s>" tag, inputString, tag)
		taggedCharacteristics = append(taggedInput, taggedString )
	}
	taggedInputsFlat := strings.Join(taggedInput, "")

	return taggedInputsFlat

}