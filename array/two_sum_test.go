package array

import (
	"fmt"
	"testing"
)

func TestTwoSumFunc(t *testing.T) {
	t.Log("Test Two Sum algorithm")
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Print("INPUT: ", nums)
	fmt.Println(", Target: ", target)
	fmt.Println("OUTPUT: ", twoSum(nums, target))
}
