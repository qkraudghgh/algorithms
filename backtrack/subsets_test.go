package backtrack

import (
	"fmt"
	"testing"
)

func TestSubsetsFunc(t *testing.T) {
	t.Log("Test Subsets algorithm")
	test := []int{1, 2, 3}

	fmt.Printf("INPUT: %v \n", test)
	fmt.Println("OUTPUT: ", subsets(test))
}
