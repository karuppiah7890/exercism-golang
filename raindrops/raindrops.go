package raindrops

import (
	"fmt"
	"strconv"
)

func appendBasedOnCondition(input string, appendContent string, condition bool) string {
	if condition {
		return fmt.Sprintf("%v%v", input, appendContent)
	}

	return input
}

// Convert converts a number to a special string
func Convert(input int) string {
	output := ""

	output = appendBasedOnCondition(output, "Pling", input%3 == 0)
	output = appendBasedOnCondition(output, "Plang", input%5 == 0)
	output = appendBasedOnCondition(output, "Plong", input%7 == 0)

	if output == "" {
		return strconv.Itoa(input)
	}

	return output
}
