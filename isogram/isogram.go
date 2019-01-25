package isogram

import "unicode"

// IsIsogram tells if the given word or phrase is an isogram or not
func IsIsogram(input string) bool {
	letterCount := make(map[rune]int)

	for _, letter := range input {
		if !unicode.IsLetter(letter) {
			continue
		}

		lowerCaseLetter := unicode.ToLower(letter)
		letterCount[lowerCaseLetter]++

		if letterCount[lowerCaseLetter] > 1 {
			return false
		}
	}

	return true
}
