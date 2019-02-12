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

	firstLetter := rune(input[0])
	if unicode.IsLetter(firstLetter) {
		letter := unicode.ToUpper(firstLetter)
		abbreviation.WriteRune(letter)
	}
	for index := 1; index < len(input); index++ {
		c := rune(input[index])
		previous := rune(input[index-1])
		if (previous == ' ' || previous == '-') && unicode.IsLetter(c) {
			letter := unicode.ToUpper(c)
			abbreviation.WriteRune(letter)
		}
	}

	return abbreviation.String()
}
