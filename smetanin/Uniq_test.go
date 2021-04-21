package smetanin_test

import (
	"testing"

	"github.com/codingunderthehood/go-course/smetanin"
)

func TestSolutionUniq(t *testing.T) {
	res0 := smetanin.SolutuionUniq([]int{1, 2, 3, 4, 5, 6, 7, 8, 8, 9, 0, 0, 0, 0, 0})
	if res0 != 10 {
		t.Fatalf("Expected 10, but got %v", res0)
	}

	res1 := smetanin.SolutuionUniq([]int{})
	if res1 != 0 {
		t.Fatalf("Expected 0, but got %v", res1)
	}

	res2 := smetanin.SolutuionUniq([]int{-1})
	if res2 != 1 {
		t.Fatalf("Expected 1, but got %v", res2)
	}
}
