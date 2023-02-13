package word

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = cases.Title(language.Dutch).String(s)
	return strings.Replace(s, " ", "", -1)
}
