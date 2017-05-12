package backtrack

import (
	"fmt"
	"testing"
)

func TestGenerateParenthesisFunc(t *testing.T) {
	t.Log("Test Generate Abbreviations algorithm")
	fmt.Println("INPUT: ", 3)
	fmt.Println("OUTPUT: ", genParenthesis(3))
}
