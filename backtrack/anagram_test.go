package backtrack

import (
	"fmt"
	"testing"
)

func TestAnagramFunc(t *testing.T) {
	t.Log("Test Anagram algorithm")
	word1 := "apple"
	word2 := "pleap"

	fmt.Printf("INPUT: %s, %s \n", word1, word2)
	fmt.Println("OUTPUT: ", anagram(word1, word2))
}
