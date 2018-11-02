package raindrops

import (
	"fmt"
	"strconv"
)

func divisibleBy(input int, number int) bool {
	if input%number == 0 {
		return true
	}
	return false
}

func appendBasedOnCondition(input string, appendContent string, condition bool) string {
	if condition {
		return fmt.Sprintf("%v%v", input, appendContent)
	}

	return input
}

// Convert converts a number to a special string
func Convert(input int) string {
	output := ""

	output = appendBasedOnCondition(output, "Pling", divisibleBy(input, 3))
	output = appendBasedOnCondition(output, "Plang", divisibleBy(input, 5))
	output = appendBasedOnCondition(output, "Plong", divisibleBy(input, 7))

	if output == "" {
		return strconv.Itoa(input)
	}

	return output
}
