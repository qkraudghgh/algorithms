package array

import (
	"testing"
	"fmt"
)

func TestLongestNonRepeatFunc(t *testing.T) {
	t.Log("Test Longest Non Repeat algorithm")
	tempString := "abcabcdefbb"
	output := longestNonRepeat(tempString)
	fmt.Println("INPUT: 'abcabcdefbb'")
	fmt.Printf("OUTPUT: %v", output)
	fmt.Println()
}
