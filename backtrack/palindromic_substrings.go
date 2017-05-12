/*
It looks like you need to be looking not for all palindromic substrings, but rather for all the ways you can divide the input string up into palindromic substrings.
(There's always at least one way, since one-character substrings are always palindromes.)

Here's the way I've done it:
*/

package backtrack

import (
	"github.com/qkraudghgh/algorithms/utils"
)

func palindromicSubstrings(s string) [][]string {
	results := [][]string{}

	palindromicBacktrack(s, &results, []string{})
	return results
}

func palindromicBacktrack(str string, results *[][]string, words []string) {
	for i := len(str); i > 0; i-- {
		sub := str[:i]
		if sub == utils.ReverseString(sub) {
			tempSub := []string{}
			for _, word := range words {
				tempSub = append(tempSub, word)
			}
			tempSub = append(tempSub, sub)
			if len(str[i:]) == 0 {
				*results = append(*results, tempSub)
			}
			palindromicBacktrack(str[i:], results, tempSub)
		}
	}
}
