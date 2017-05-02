package array

import (
	"fmt"
	"testing"
)

func TestPlusOneFunc(t *testing.T) {
	t.Log("Test Plus One algorithm")
	nums := []int{1, 2, 3, 4, 5, 11, 12, 13}
	fmt.Println("INPUT: ", nums)
	fmt.Println("OUTPUT: ", plusOne(nums))
}
