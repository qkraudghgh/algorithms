package backtrack

import (
	"fmt"
	"testing"
)

func TestCombinationSumFunc(t *testing.T) {
	t.Log("Test Combinations Sum algorithm")
	tempSlice := []int{2, 3, 6, 7}
	target := 6
	fmt.Printf("INPUT: %v, Target: %d", tempSlice, target)
	fmt.Println()
	fmt.Print("OUTPUT:", combinationSum(tempSlice, target))
}
