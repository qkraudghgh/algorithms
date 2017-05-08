package array

import (
	"fmt"
	"testing"
)

func TestJosepheusFunc(t *testing.T) {
	t.Log("Test Circular Counter algorithm")
	temp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("INPUT: ", temp)
	fmt.Print("OUTPUT: ")
	josepheus(temp, 3)
	fmt.Println()
}
