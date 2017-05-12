/*
Given a digit string, return all possible letter
combinations that the number could represent.

Input:Digit string "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
*/

package backtrack

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	kmaps := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	ans := []string{""}

	for _, num := range digits {
		tmp := []string{}
		tempNum := string(num) // byte to string
		for _, an := range ans {
			for _, char := range kmaps[tempNum] {
				tmp = append(tmp, an+string(char))
			}
		}
		ans = tmp
	}
	return ans
}
