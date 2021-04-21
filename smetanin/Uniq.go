package smetanin

import "fmt"

func SolutuionUniq(A []int) int {
	m := make(map[int]bool)
	for _, element := range A {
		m[element] = true
	}

	fmt.Println(m)

	return len(m)
}
