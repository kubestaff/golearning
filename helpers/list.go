package helpers

import (
	"fmt"
	"strings"
)

func FormatStringstoTags(inputStrings []string, tag string) string {

	formattedInputs := []string{}
	for _, inputString := range inputStrings {
		formattedInputString := fmt.Sprintf("<%s>%s</%s>", tag, inputString, tag)
		formattedInputs = append(formattedInputs, formattedInputString)
	}
	joinedCharacteristics := strings.Join(formattedInputs, "")

	return joinedCharacteristics
}
