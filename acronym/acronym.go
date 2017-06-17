package acronym

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

const testVersion = 3

//Abbreviate generates the acronym from the provided string. This
func Abbreviate(phrase string) string {
	splitter, _ := regexp.Compile(`[\s-]`)
	words := splitter.Split(phrase, -1)

	var acronym string
	for _, word := range words {
		firstRune, _ := utf8.DecodeRuneInString(word)

		acronym = acronym + strings.ToUpper(string(firstRune))
	}
	return acronym
}
