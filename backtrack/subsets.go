/*
Given a set of distinct integers, nums, return all possible subsets.

Note: The solution set must not contain duplicate subsets.

For example,
If nums = [1,2,3], a solution is:

[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]

O(2**n)
*/

package backtrack

import "github.com/qkraudghgh/algorithms/utils"

func subsets(nums []int) [][]int {
	res := [][]int{}
	subsetBacktrack(&res, nums, []int{}, 0)
	return res
}

func subsetBacktrack(res *[][]int, nums []int, stack []int, pos int) {
	if pos == len(nums) {
		*res = append(*res, stack)
	} else {
		stack = append(stack, nums[pos])
		subsetBacktrack(res, nums, stack, pos+1)
		stack = utils.Pop(stack)
		subsetBacktrack(res, nums, stack, pos+1)
	}
}
