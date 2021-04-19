package smetanin_test

import (
	"reflect"
	"testing"

	"github.com/codingunderthehood/go-course/smetanin"
)

func TestSolutionSquareGenerator(t *testing.T) {
	res0 := smetanin.SolutionSquareGenerator(10, 2)
	if !reflect.DeepEqual(res0, []int{100, 121}) {
		t.Fatalf("[10, 11] should become [100, 121], but got %v", res0)
	}

	res1 := smetanin.SolutionSquareGenerator(-2, 5)
	if !reflect.DeepEqual(res1, []int{1, 4}) {
		t.Fatalf("[-2, -1, 0, 1, 2] should become [1, 4], but got %v", res1)
	}

	res2 := smetanin.SolutionSquareGenerator(0, 3)
	if !reflect.DeepEqual(res2, []int{1, 4}) {
		t.Fatalf("[0, 1, 2] should become [1, 4], but got %v", res2)
	}

	res3 := smetanin.SolutionSquareGenerator(0, 0)
	if !reflect.DeepEqual(res3, []int{}) {
		t.Fatalf("Should return empty slice for 0 elements")
	}

	res4 := smetanin.SolutionSquareGenerator(0, -1)
	if !reflect.DeepEqual(res4, []int{}) {
		t.Fatalf("Should return empty slice for negative number of elements")
	}
}
