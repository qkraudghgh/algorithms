package backtrack

import (
	"fmt"
	"testing"
)

func TestLetterCombinationFunc(t *testing.T) {
	t.Log("Test Letter Combination algorithm")
	digitNum := "23"

	fmt.Printf("INPUT: %s \n", digitNum)
	fmt.Println("OUTPUT: ", letterCombinations(digitNum))
}
