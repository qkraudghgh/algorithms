package backtrack

import (
	"fmt"
	"testing"
)

func TestPalindromicSubstringsFunc(t *testing.T) {
	t.Log("Test Palindromic Substrings algorithm")
	word := "banana"

	fmt.Printf("INPUT: %v \n", word)
	fmt.Println("OUTPUT: ", palindromicSubstrings(word))
}
