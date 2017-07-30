package acronym

import (
	"regexp"
	"strings"
)

const testVersion = 3

// Abbreviate returns an abbreviation for the supplied input
func Abbreviate(input string) string {
	var result string
	re := regexp.MustCompile("([[:alpha:]])[[:alpha:]]+")

	matches := re.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		result += match[1]
	}

	return strings.ToUpper(result)
}
