package scrabble

import "unicode"

func getScore(letter rune) int {
	switch letter {
	case 'a':
		fallthrough
	case 'e':
		fallthrough
	case 'i':
		fallthrough
	case 'o':
		fallthrough
	case 'u':
		fallthrough
	case 'l':
		fallthrough
	case 'n':
		fallthrough
	case 'r':
		fallthrough
	case 's':
		fallthrough
	case 't':
		return 1
	case 'd':
		fallthrough
	case 'g':
		return 2
	case 'b':
		fallthrough
	case 'c':
		fallthrough
	case 'm':
		fallthrough
	case 'p':
		return 3
	case 'f':
		fallthrough
	case 'h':
		fallthrough
	case 'v':
		fallthrough
	case 'w':
		fallthrough
	case 'y':
		return 4
	case 'k':
		return 5
	case 'j':
		fallthrough
	case 'x':
		return 8
	case 'q':
		fallthrough
	case 'z':
		return 10
	default:
		return 0
	}
}

// Score finds the scrabble score for a word
func Score(word string) int {
	score := 0

	for _, letter := range word {
		lowerCaseLetter := unicode.ToLower(letter)
		score += getScore(lowerCaseLetter)
	}

	return score
}
