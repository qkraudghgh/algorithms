/*
Given a pattern and a string str,
find if str follows the same pattern.

Here follow means a full match, such that there is a bijection between
a letter in pattern and a non-empty substring in str.

Examples:
pattern = "abab", str = "redblueredblue" should return true.
pattern = "aaaa", str = "asdasdasdasd" should return true.
pattern = "aabb", str = "xyzabcxzyabc" should return false.
Notes:
You may assume both pattern and str contains only lowercase letters.
*/

package backtrack

import "github.com/qkraudghgh/algorithms/utils"

func patternMatch(pattern string, str string) bool {
	dic := make(map[byte]string)
	return patternMatcher(pattern, str, &dic)
}

func patternMatcher(pattern string, str string, dic *map[byte]string) bool {
	tempdic := *dic

	// if you want to show map flow, delete the comment below
	// fmt.Println(tempdic)
	if len(pattern) == 0 && len(str) > 0 {
		return false
	}
	if len(pattern) == 0 && len(str) == 0 {
		return true
	}
	for i := 1; i < len(str)-len(pattern)+2; i++ {
		if _, ok := tempdic[pattern[0]]; !ok {
			if !utils.FindStrByByteInMap(str[:i], tempdic) {
				tempdic[pattern[0]] = str[:i]
				if patternMatcher(pattern[1:], str[i:], dic) {
					return true
				}
				delete(tempdic, pattern[0])
			}
		} else if _, ok := tempdic[pattern[0]]; ok {
			if tempdic[pattern[0]] == str[:i] {
				if patternMatcher(pattern[1:], str[i:], dic) {
					return true
				}
			}
		}
	}
	return false
}
