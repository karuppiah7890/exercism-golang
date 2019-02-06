// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate abbreviates a given sentence or phrase
func Abbreviate(input string) string {
	abbreviation := strings.Builder{}

	for index, c := range input {
		if index > 0 {
			previous := rune(input[index-1])
			if (previous == ' ' || previous == '-') && unicode.IsLetter(c) {
				letter := unicode.ToUpper(c)
				abbreviation.WriteRune(letter)
			}
		} else if unicode.IsLetter(c) {
			letter := unicode.ToUpper(c)
			abbreviation.WriteRune(letter)
		}
	}

	return abbreviation.String()
}
