package array

import (
	"fmt"
	"testing"
)

func TestRotateSliceFunc(t *testing.T) {
	t.Log("Test Rotate Slice algorithm")
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Print("INPUT: ")
	fmt.Println(nums)
	rotate(nums, 3)
	fmt.Print("OUTPUT: ")
	fmt.Println(nums)
}
