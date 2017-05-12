package backtrack

import (
	"fmt"
	"testing"
)

func TestAllPermsFunc(t *testing.T) {
	t.Log("Test All Perms algorithm")
	word := "abc"

	fmt.Println("INPUT: ", word)
	fmt.Println("OUTPUT: ", allPerms(word))
}
