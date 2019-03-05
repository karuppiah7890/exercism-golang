// Package luhn contains Valid() function to check the validity of a number
// as per the luhn formula
package luhn

import (
	"fmt"
	"strings"
	"unicode"
)

func getDigits(number string) ([]int, error) {
	digits := make([]int, 0, len(number))

	for _, digit := range number {
		if !unicode.IsDigit(digit) {
			return nil, fmt.Errorf("non-digit characters present")
		}

		digits = append(digits, int(digit-'0'))
	}

	return digits, nil
}

func doubleDigit(d int) int {
	doubledValue := d * 2

	if doubledValue >= 10 {
		return doubledValue - 9
	}

	return doubledValue
}

func sum(digits []int) int {
	sum := 0
	length := len(digits)
	lengthIsOdd := length % 2

	for index, digit := range digits {
		if index%2 == lengthIsOdd {
			value := doubleDigit(digit)
			sum += value
			continue
		}
		sum += digit
	}
	return sum
}

// Valid checks if a number is valid or not as per luhn's formula
func Valid(number string) bool {
	number = strings.Replace(number, " ", "", -1)

	if len(number) <= 1 {
		return false
	}

	digits, err := getDigits(number)

	if err != nil {
		return false
	}

	sumOfDigits := sum(digits)

	if sumOfDigits%10 != 0 {
		return false
	}

	return true
}
