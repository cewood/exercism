package pangram

import "strings"

const testVersion = 1

// IsPangram checks if a sentence is a pangram
func IsPangram(input string) bool {
	alphabet := []string{
		"a", "b", "c", "d", "e",
		"f", "g", "h", "i", "j",
		"k", "l", "m", "n", "o",
		"p", "q", "r", "s", "t",
		"u", "v", "w", "x", "y", "z",
	}

	for _, letter := range alphabet {
		if !strings.Contains(strings.ToLower(input), letter) {
			return false
		}
	}

	return true
}
