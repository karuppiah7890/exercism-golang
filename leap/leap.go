package leap

const testVersion = 3

// IsLeapYear takes year as parameter and returns true if it's leap year or false otherwise
func IsLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	}
	if year%100 != 0 {
		return true
	}
	if year%400 != 0 {
		return false
	}
	return true
}
