package backtrack

import (
	"fmt"
	"testing"
)

func TestWordSearchFunc(t *testing.T) {
	t.Log("Test Word Search algorithm")
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}

	words := []string{"oath", "pea", "eat", "rain"}

	fmt.Printf("Board: %v, Words: %v \n", board, words)
	fmt.Println("OUTPUT: ", findWords(board, words))
}
