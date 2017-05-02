package array

import (
	"fmt"
	"testing"
)

func TestMissingRangesFunc(t *testing.T) {
	t.Log("Test Missing Ranges algorithm")
	nums := []int{3, 5, 10, 11, 12, 15, 19}
	fmt.Println("INPUT: ", nums)
	fmt.Println("OUTPUT: ", missingRanges(nums, 0, 20))
}
