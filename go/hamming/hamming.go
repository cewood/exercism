package hamming

import "fmt"

const testVersion = 6

// Distance returns the Hamming distance for two inputs
func Distance(a, b string) (int, error) {
	var hammingDistance int

	if len(a) != len(b) {
		return -1, fmt.Errorf("Supplied input lengths don't match")
	}

	for i := 0; i <= len(a)-1; i++ {
		if a[i] != b[i] {
			hammingDistance++
		}
	}
	return hammingDistance, nil
}
