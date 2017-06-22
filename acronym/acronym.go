package acronym

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

const testVersion = 3

//Abbreviate generates the acronym from the provided string. This
func Abbreviate(phrase string) string {
	wordDelimiter, _ := regexp.Compile(`[\s-]`)

	var acronym string
	for _, word := range wordDelimiter.Split(phrase, -1) {
		firstRune, _ := utf8.DecodeRuneInString(word)

		acronym = acronym + strings.ToUpper(string(firstRune))
	}
	return acronym
}
