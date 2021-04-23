package smetanin

func SolutuionUniq(A []int) int {
	m := make(map[int]bool)
	for _, element := range A {
		m[element] = true
	}

	return len(m)
}
