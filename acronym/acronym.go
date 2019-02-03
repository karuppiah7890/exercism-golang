// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"fmt"
	"strings"
)

// Abbreviate abbreviates a given sentence or phrase
func Abbreviate(input string) string {
	words := strings.Split(input, " ")
	abbreviation := ""

	for _, word := range words {
		abbreviation = fmt.Sprintf("%v%c", abbreviation, rune(word[0]))
	}

	return abbreviation
}
