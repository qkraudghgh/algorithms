package array

import (
	"fmt"
	"testing"
)

func TestSummaryRangesFunc(t *testing.T) {
	t.Log("Test Summary Ranges algorithm")
	nums := []int{0, 1, 2, 4, 5, 7}
	fmt.Print("INPUT: ")
	fmt.Println(nums)
	fmt.Print("OUTPUT: ")
	fmt.Println(summaryRanges(nums))
}
