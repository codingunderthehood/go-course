package smetanin

// It is worth to note that we hope that user won't input some data that will lead to an int overflow
func SolutionSquareGenerator(start int, n int) []int {
	// return empty on negative and zero number of elements
	if n <= 0 {
		return []int{}
	}
	// skip non natural numbers and zero
	if start <= 0 {
		n += start - 1
		start = 1
	}

	result := make([]int, n)
	curr := start
	for i := 0; i < n; i++ {
		result[i] = curr * curr
		curr += 1
	}
	return result
}
