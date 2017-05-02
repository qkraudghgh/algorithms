package array

import (
	"fmt"
	"testing"
)

func TestGarageFunc(t *testing.T) {
	t.Log("Test Garage algorithm")
	initial := []int{1, 2, 3, 0, 4}
	final := []int{0, 3, 2, 1, 4}
	fmt.Println("INITIAL: ", initial)
	fmt.Println("FINAL: ", final)
	fmt.Println(garage(initial, final))
}
