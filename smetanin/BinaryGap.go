package smetanin

// It is worth to note that we hope that user won't input some data that will lead to an int overflow
func SolutionBinaryGap(N int) int {
	if N < 0 {
		return 0
	}
	if N == 0 {
		return 1
	}

	distance := 0
	result := 0

	for N > 0 {
		if N&1 == 0 {
			N = N >> 1
			distance += 1
		} else {
			if distance > result {
				result = distance
			}
			distance = 0
			N = N >> 1
		}
	}

	return result
}
