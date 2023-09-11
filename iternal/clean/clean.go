package clean

import (
	"regexp"
	"strings"
)

func Name(prodName string) string {
	prodName = strings.ReplaceAll(prodName, "\n", "")
	prodName = strings.ReplaceAll(prodName, "\"", "")
	prodName = strings.ReplaceAll(prodName, `"\`, "")
	return prodName
}

func Number(prodNumber string) string {
	regex := regexp.MustCompile("\\s+")
	prodNumber = strings.ReplaceAll(prodNumber, "\n", "")

	prodNumber = regex.ReplaceAllString(prodNumber, " ")
	return prodNumber
}
