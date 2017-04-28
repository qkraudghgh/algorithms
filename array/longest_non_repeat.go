/*
Given a string, find the length of the longest substring
without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3.
Note that the answer must be a substring,
"pwke" is a subsequence and not a substring.
*/

package array

func longestNonRepeat(s string) int {
	var start, maxlen int
	used_char := make(map[rune]int)
	for index, char := range s {
		_, ok := used_char[char]
		if ok && start <= used_char[char] {
			start = used_char[char] + 1
		} else {
			temp := index - start + 1
			if maxlen < temp {
				maxlen = temp
			}
		}
		used_char[char] = index
	}
	return maxlen
}
