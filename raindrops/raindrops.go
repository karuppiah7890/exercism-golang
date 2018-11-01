package raindrops

import "fmt"

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

	divisibleByThree := divisibleBy(input, 3)
	divisibleByFive := divisibleBy(input, 5)
	divisibleBySeven := divisibleBy(input, 7)

	if !divisibleByThree && !divisibleByFive && !divisibleBySeven {
		return fmt.Sprintf("%v", input)
	}

	output := ""

	output = appendBasedOnCondition(output, "Pling", divisibleByThree)
	output = appendBasedOnCondition(output, "Plang", divisibleByFive)
	output = appendBasedOnCondition(output, "Plong", divisibleBySeven)

	return output
}
