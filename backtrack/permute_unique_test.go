package backtrack

import (
	"fmt"
	"testing"
)

func TestPermuteUniqueFunc(t *testing.T) {
	t.Log("Test Permute Unique algorithm")
	test := []int{1, 1, 2}

	fmt.Printf("INPUT: %v \n", test)
	fmt.Println("OUTPUT: ", permuteUnique(test))
}
