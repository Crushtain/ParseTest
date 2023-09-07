package iternal

import "strings"

func Clean(prodName string) string {
	prodName = strings.ReplaceAll(prodName, "\n", "")
	prodName = strings.ReplaceAll(prodName, "\"", "")
	prodName = strings.ReplaceAll(prodName, `"\`, "")
	return prodName
}
