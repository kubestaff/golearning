package helper

import (
	"fmt"
	"strings"
)

func FormatStringsToTags(inputStrings []string, tag string) string {

	formattedInput := []string{}
	for_, inputString := range inputinputStrings {
		formattedInputStrings := fmt.Sprintf("<%s>%s</%s>" tag, inputString, tag)
		formattedInput = append(formattedInputs, formattedInputString )
	}
	joinedCharacteristics := strings.Join(formattedInputs, "")

	return joinedCharacteristics

}