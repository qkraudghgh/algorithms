/*
Given n pairs of parentheses, write a function to generate
all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/

package backtrack

func genParenthesis(n int) []string {
	res := []string{}
	addPair(&res, "", n, 0)
	return res
}

func addPair(res *[]string, s string, left int, right int) {
	if left == 0 && right == 0 {
		*res = append(*res, s)
		return
	}
	if right > 0 {
		addPair(res, s+")", left, right-1)
	}
	if left > 0 {
		addPair(res, s+"(", left-1, right+1)
	}
}
