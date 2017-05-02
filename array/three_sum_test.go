package array

import (
	"fmt"
	"testing"
)

func TestThreeSumFunc(t *testing.T) {
	t.Log("Test Three Sum algorithm")
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println("INPUT: ", nums)
	fmt.Println("OUTPUT: ", threeSum(nums))
}
