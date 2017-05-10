package backtrack

import (
	"fmt"
	"testing"
)

func TestArraySumCombinationsFunc(t *testing.T) {
	t.Log("Test Array Sum Combinations algorithm")
	tempSlice := []int{}
	target := 7
	fmt.Println("INPUT: [], Target:", target)
	fmt.Print("OUTPUT: ")
	sumCombinationsBacktrack(&tempSlice, target)
	fmt.Println()
}
