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
	betweenWords := true
	for _, c := range input {
		if c == ' ' || c == '-' {
			betweenWords = true
			continue
		}
		if betweenWords {
			abbreviation.WriteRune(unicode.ToUpper(c))
			betweenWords = false
		}
	}

	return abbreviation.String()
}
