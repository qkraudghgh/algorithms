/*
Given a set of candidate numbers (C) (without duplicates) and a target number
(T), find all unique combinations in C where the candidate numbers sums to T.

The same repeated number may be chosen from C unlimited number of times.

Note:
All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.
For example, given candidate set [2, 3, 6, 7] and target 7,
A solution set is:
[
  [7],
  [2, 2, 3]
]
*/

package backtrack

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)
	dfs(candidates, target, 0, []int{}, &res)
	return res
}

func dfs(nums []int, target int, index int, path []int, res *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}
	for i := index; i < len(nums); i++ {
		dfs(nums, target-nums[i], i, append(path, nums[i]), res)
	}
}
