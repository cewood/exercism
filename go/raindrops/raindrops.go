package raindrops

import "fmt"

const testVersion = 3

// Convert returns some output based on the divisors of a number
func Convert(i int) string {
	var response string

	for j := i; j > 0; j-- {
		if i%j == 0 {
			if j == 3 {
				response = "Pling" + response
			} else if j == 5 {
				response = "Plang" + response
			} else if j == 7 {
				response = "Plong" + response
			}
		}
	}

	if response == "" {
		response = fmt.Sprintf("%v", i)
	}

	return response
}
