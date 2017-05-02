package array

import (
	"fmt"
	"testing"
)

func TestMergeIntervalsFunc(t *testing.T) {
	t.Log("Test Merge Intervals algorithm")
	temps := intervals{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	fmt.Println("INPUT: ", temps)
	fmt.Println("OUTPUT: ", merge(temps))
}
