package backtrack

import (
	"fmt"
	"testing"
)

func TestThreeSumFunc(t *testing.T) {
	t.Log("Test Generate Abbreviations algorithm")
	word := "word"
	fmt.Println("INPUT: ", word)
	fmt.Println("OUTPUT: ", generateAbbreviations(word))
}
