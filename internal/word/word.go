package word

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
	"unicode"
)

func ToUpper(str string) string {
	return strings.ToUpper(str)
}

func ToLower(str string) string {
	return strings.ToLower(str)
}

func UnderscoreToUpperCamelCase(str string) string {
	str = strings.Replace(str, "_", " ", -1)
	str = cases.Title(language.Und).String(str)
	str = strings.Replace(str, " ", "", -1)
	return str
}

func UnderscoreToLowerCamelCase(str string) string {
	str = UnderscoreToUpperCamelCase(str)
	return string(unicode.ToLower(rune(str[0]))) + str[1:]
}

func CamelCaseToUnderscore(str string) string {
	// Use govalidator
	var output []rune
	for i, r := range str {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}

	return string(output)
}
