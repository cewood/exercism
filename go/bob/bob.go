package bob

import (
	"regexp"
	"strings"
)

const testVersion = 3

const responseQuestion = "Sure."
const responseYelling = "Whoa, chill out!"
const responseSilence = "Fine. Be that way!"
const responseDefault = "Whatever."

// Hey returns a specific response depending on the input
func Hey(input string) string {
	questionRe := regexp.MustCompile("[?]+[\t\n\f\r ]*$")
	digitsRe := regexp.MustCompile("[[:digit:]]+")
	silenceRe := regexp.MustCompile("^[\n\r\t ]+$")
	normalRe := regexp.MustCompile("[[:alpha:]]+")

	if input == strings.ToUpper(input) &&
		len(normalRe.FindAllStringSubmatch(input, -1)) > 0 {
		return responseYelling
	} else if len(questionRe.FindAllStringSubmatch(input, -1)) > 0 {
		return responseQuestion
	} else if len(silenceRe.FindAllStringSubmatch(input, -1)) > 0 &&
		len(digitsRe.FindAllStringSubmatch(input, -1)) == 0 || input == "" {
		return responseSilence
	}

	return responseDefault
}
