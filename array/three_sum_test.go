package array

import (
	"fmt"
	"testing"
)

func TestThreeSumFunc(t *testing.T) {
	t.Log("Test Three Sum algorithm")
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Print("INPUT: ")
	fmt.Println(nums)
	fmt.Print("OUTPUT: ")
	fmt.Println(threeSum(nums))
}
