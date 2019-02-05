// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate abbreviates a given sentence or phrase
func Abbreviate(input string) string {
	input = strings.Replace(input, "-", " ", -1)
	words := strings.Split(input, " ")
	abbreviation := make([]rune, 0)

	for _, word := range words {
		if len(word) == 0 {
			continue
		}
		letter := unicode.ToUpper(rune(word[0]))
		abbreviation = append(abbreviation, letter)
	}

	return string(abbreviation)
}
