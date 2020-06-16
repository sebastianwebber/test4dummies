package style

import (
	"errors"
	"strings"
	"unicode"
)

// Sentence converts the input test,
// adding upper case to the first letter of each sentence.
func Sentence(input string) (string, error) {

	// For test purposes only
	if input == "" {
		return "", errors.New("Could not process empty strings")
	}

	newSentence := true
	return strings.Map(
		func(r rune) rune {
			if newSentence && r != ' ' {
				newSentence = false
				return unicode.ToTitle(r)
			}

			if r == '.' {
				newSentence = true
			}

			return r
		}, input), nil

}
