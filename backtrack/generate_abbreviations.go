/*
given input word, return the list of abbreviations.
ex)
word => [1ord, w1rd, wo1d, w2d, 3d, w3 ... etc]
*/
package backtrack

import (
	"bytes"
	"strconv"
)

func generateAbbreviations(word string) []string {
	result := []string{}
	wordBacktrack(&result, word, 0, 0, "")
	return result
}

func wordBacktrack(result *[]string, word string, pos int, count int, cur string) {
	if pos == len(word) {
		if count > 0 {
			cur += strconv.Itoa(count)
		}
		// because Go is passed by value
		*result = append(*result, cur)
		return
	}

	if count > 0 {
		tempBuff := bytes.NewBufferString(cur + strconv.Itoa(count))
		tempBuff.WriteByte(word[pos])
		wordBacktrack(result, word, pos+1, 0, tempBuff.String())
	} else {
		tempBuff := bytes.NewBufferString(cur)
		tempBuff.WriteByte(word[pos])
		wordBacktrack(result, word, pos+1, 0, tempBuff.String())
	}
	wordBacktrack(result, word, pos+1, count+1, cur)
}
