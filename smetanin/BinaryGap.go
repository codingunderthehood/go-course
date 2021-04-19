package smetanin

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

	if distance > result {
		result = distance
	}

	return result
}
