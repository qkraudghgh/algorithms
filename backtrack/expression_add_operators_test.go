package backtrack

import (
	"fmt"
	"testing"
)

func TestExpressionAddOperatorsFunc(t *testing.T) {
	t.Log("Test Expression Add Operators algorithm")
	num := "123"
	target := 6
	// "123", 6 -> ["1+2+3", "1*2*3"]
	fmt.Printf("INPUT: %s, Target: %d \n", num, target)
	fmt.Println("OUTPUT:", addOperator(num, target))

	// "232", 8 -> ["2*3+2", "2+3*2"]
	num = "232"
	target = 8
	fmt.Printf("INPUT: %s, Target: %d \n", num, target)
	fmt.Println("OUTPUT:", addOperator(num, target))

	// "123045", 3
	num = "123045"
	target = 3
	fmt.Printf("INPUT: %s, Target: %d \n", num, target)
	fmt.Println("OUTPUT:", addOperator(num, target))
}
