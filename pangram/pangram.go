//Package pangram offers methods for identifying pangram phrases.
package pangram

import "strings"

const testVersion = 1

//IsPangram identifies if a phrase is a pangram - it contains all the letters of the alphabet.
func IsPangram(phrase string) bool {
	lowerCasePhrase := strings.ToLower(phrase)

	for _, letter := range alphabetLetters() {
		if !strings.Contains(lowerCasePhrase, string(letter)) {
			return false
		}
	}
	return true
}

func alphabetLetters() []rune {
	var letters []rune
	for i := 'a'; i < 'z'; i++ {
		letters = append(letters, i)
	}
	return letters
}
