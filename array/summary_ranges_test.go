package array

import (
	"fmt"
	"testing"
)

func TestSummaryRangesFunc(t *testing.T) {
	t.Log("Test Summary Ranges algorithm")
	nums := []int{0, 1, 2, 4, 5, 7}
	fmt.Println("INPUT: ", nums)
	fmt.Println("OUTPUT: ", summaryRanges(nums))
}
