//Package bob simulates a lackadaisical teenager
package bob

import (
	"regexp"
	"strings"
)

const testVersion = 3

//Hey function gives the answers of a lackadaisical teenager to a phrase
func Hey(phrase string) string {
	if yell(phrase) {
		return "Whoa, chill out!"
	}
	if question(phrase) {
		return "Sure."
	}
	if silence(phrase) {
		return "Fine. Be that way!"
	}
	return "Whatever."
}

func yell(phrase string) bool {
	everythingButLettersFinder, _ := regexp.Compile(`[^\w]|[0-9]`)

	onlyCharacters := everythingButLettersFinder.ReplaceAllString(phrase, "")

	return onlyCharacters == strings.ToUpper(onlyCharacters) && len(onlyCharacters) > 0
}

func silence(phrase string) bool {
	nonSpaceFinder, _ := regexp.Compile(`[^\s]`)
	return !nonSpaceFinder.MatchString(phrase)
}

func question(phrase string) bool {
	questionMarkFinder, _ := regexp.Compile("[?][\\s]*$")
	return questionMarkFinder.MatchString(phrase)

}
