// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "unicode"

// Hey returns bob's response to a remark
func Hey(remark string) string {
	question := false
	numberOfLetters, numberOfUpperCaseLetters := 0, 0
	for _, letter := range remark {
		if !unicode.IsLetter(letter) {
			if letter == '?' {
				question = true
			}
			continue
		}

		numberOfLetters++

		if unicode.IsUpper(letter) {
			numberOfUpperCaseLetters++
		}
	}

	shouting := false

	if numberOfLetters > 0 && numberOfUpperCaseLetters == numberOfLetters {
		shouting = true
	}

	if shouting && question {
		return "Calm down, I know what I'm doing!"
	}

	if shouting {
		return "Whoa, chill out!"
	}

	if question {
		return "Sure."
	}

	return "Whatever."
}
