package array

import (
	"fmt"
	"testing"
)

func TestTwoSumFunc(t *testing.T) {
	t.Log("Test Two Sum algorithm")
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Print("INPUT: ")
	fmt.Print(nums)
	fmt.Printf(", Target: %v", target)
	fmt.Println()
	fmt.Print("OUTPUT: ")
	fmt.Println(twoSum(nums, target))
}
