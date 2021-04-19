package smetanin_test

import (
	"testing"

	"github.com/codingunderthehood/go-course/smetanin"
)

func TestSolutionBinaryGap(t *testing.T) {
	res0 := smetanin.SolutionBinaryGap(4)
	if res0 != 2 {
		t.Fatalf("Expected 2, but got %v", res0)
	}

	res1 := smetanin.SolutionBinaryGap(17)
	if res1 != 3 {
		t.Fatalf("Expected 3, but got %v", res1)
	}

	res2 := smetanin.SolutionBinaryGap(-1)
	if res2 != 0 {
		t.Fatalf("Expected 0, but got %v", res2)
	}

	res3 := smetanin.SolutionBinaryGap(3)
	if res3 != 0 {
		t.Fatalf("Expected 0, but got %v", res3)
	}
}
