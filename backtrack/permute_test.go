package backtrack

import (
	"fmt"
	"testing"
)

func TestPermuteFunc(t *testing.T) {
	t.Log("Test Permute algorithm")
	test := []int{1, 2, 3}

	fmt.Printf("INPUT: %v \n", test)
	fmt.Println("OUTPUT: ", permute(test))
}
