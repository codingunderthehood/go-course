package smetanin

func SolutuionArrayShift(A []int, K int) {
	if len(A) == 0 {
		return
	}
	k := K % len(A)
	if k == 0 {
		return
	}
	shifted := make([]int, len(A))

	shift := len(A) - k

	// fill out first part – elements that moved to the head
	for index, element := range A[shift:] {
		shifted[index] = element
	}

	// fill out second part - shifted elements
	for index, element := range A[:shift] {
		shifted[k+index] = element
	}

	copy(A, shifted)
}
