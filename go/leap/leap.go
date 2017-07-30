package leap

const testVersion = 3

// IsLeapYear determines if a year is a Gregorian leap year
func IsLeapYear(i int) bool {
	var result bool

	if i%4 == 0 &&
		i%100 != 0 ||
		i%400 == 0 {
		result == true
	}
	return result
}
