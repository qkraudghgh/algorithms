package backtrack

import (
	"fmt"
	"testing"
)

func TestFactorCombinationsFunc(t *testing.T) {
	t.Log("Test Factor Combinations algorithm")
	fmt.Println("INPUT: ", 32)
	fmt.Println("OUTPUT: ", iterativeGetFactors(32))
}
