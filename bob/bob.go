// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "unicode"

// Hey returns bob's response to a remark
func Hey(remark string) string {
	shouting := true
	for _, letter := range remark {
		if !unicode.IsLetter(letter) {
			continue
		}

		if unicode.IsLower(letter) {
			shouting = false
		}
	}

	if shouting {
		return "Whoa, chill out!"
	}

	return "Whatever."
}
