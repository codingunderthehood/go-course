package smetanin_test

import (
	"testing"

	"github.com/codingunderthehood/go-course/smetanin"
)

func TestSolutionArrayShift(t *testing.T) {
	res0 := []int{1, 2, 3, 4}
	smetanin.SolutuionArrayShift(res0, 2)
	exp0 := []int{3, 4, 1, 2}
	if !equal(res0, exp0) {
		t.Fatalf("Expected %v, but got %v", exp0, res0)
	}

	res1 := []int{1, 2, 3, 4}
	smetanin.SolutuionArrayShift(res1, 4)
	exp1 := []int{1, 2, 3, 4}
	if !equal(res1, exp1) {
		t.Fatalf("Expected %v, but got %v", exp1, res1)
	}

	res2 := []int{1, 2, 3, 4}
	smetanin.SolutuionArrayShift(res2, 1)
	exp2 := []int{4, 1, 2, 3}
	if !equal(res2, exp2) {
		t.Fatalf("Expected %v, but got %v", exp2, res2)
	}

	res3 := []int{1, 2, 3, 4}
	smetanin.SolutuionArrayShift(res3, 5)
	exp3 := []int{4, 1, 2, 3}
	if !equal(res3, exp3) {
		t.Fatalf("Expected %v, but got %v", exp3, res3)
	}

	res4 := []int{}
	smetanin.SolutuionArrayShift(res4, 1)
	exp4 := []int{}
	if !equal(res4, exp4) {
		t.Fatalf("Expected %v, but got %v", exp4, res4)
	}

	res5 := []int{1, 2, 3}
	smetanin.SolutuionArrayShift(res5, 8)
	exp5 := []int{2, 3, 1}
	if !equal(res5, exp5) {
		t.Fatalf("Expected %v, but got %v", exp5, res5)
	}
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
