package array

import (
	"fmt"
	"testing"
)

func TestLongestNonRepeatFunc(t *testing.T) {
	t.Log("Test Longest Non Repeat algorithm")
	tempString := "abcabcdefbb"
	output := longestNonRepeat(tempString)
	fmt.Println("INPUT: ", tempString)
	fmt.Println("OUTPUT: ", output)
}
