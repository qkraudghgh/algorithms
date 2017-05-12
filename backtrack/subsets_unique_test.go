package backtrack

import (
	"fmt"
	"testing"
)

func TestSubsetsUniqueFunc(t *testing.T) {
	t.Log("Test Subsets algorithm")
	test := []int{1, 2, 2}

	fmt.Printf("INPUT: %v \n", test)
	fmt.Println("OUTPUT: ", subsetsUnique(test))
}
