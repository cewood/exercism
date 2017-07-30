package accumulate

const testVersion = 1

// Accumulate peroforms an operation on each element of a collection
func Accumulate(input []string, operation func(string) string) []string {
	for i := 0; i < len(input); i++ {
		input[i] = operation(input[i])
	}

	return input
}
