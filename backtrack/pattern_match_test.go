package backtrack

import (
	"fmt"
	"testing"
)

func TestPatternMatchFunc(t *testing.T) {
	t.Log("Test Pattern Match algorithm")
	pattern1 := "abab"
	string1 := "redblueredblue"
	pattern2 := "aaaa"
	string2 := "asdasdasdasd"
	pattern3 := "aabb"
	string3 := "xyzabcxzyabc"

	// true is correct
	fmt.Printf("pattern = '%s', str = '%s' should return %t. \n", pattern1, string1, patternMatch(pattern1, string1))
	fmt.Printf("pattern = '%s', str = '%s' should return %t. \n", pattern2, string2, patternMatch(pattern2, string2))

	// false is correct
	fmt.Printf("pattern = '%s', str = '%s' should return %t. \n", pattern3, string3, patternMatch(pattern3, string3))
}
