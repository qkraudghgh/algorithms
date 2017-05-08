package array

import (
	"fmt"
	"testing"
)

func TestFlattenFunc(t *testing.T) {
	t.Log("Test Flatten algorithm")
	target := any{2, 1, any{3, any{4, 5}, 6}, 7, any{8}}
	fmt.Println("INPUT: ", target)
	fmt.Println("OUTPUT: ", flatten(target, any{}))
}
