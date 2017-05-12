/*
Given a collection of integers that might contain duplicates, nums,
return all possible subsets.

Note: The solution set must not contain duplicate subsets.

For example,
If nums = [1,2,2], a solution is:

[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]

*/

package backtrack

import (
	"github.com/qkraudghgh/algorithms/utils"
	"reflect"
)

func subsetsUnique(nums []int) [][]int {
	res := [][]int{}
	subsetUniqueBacktrack(&res, nums, []int{}, 0)
	return res
}

func subsetUniqueBacktrack(res *[][]int, nums []int, stack []int, pos int) {
	if pos == len(nums) {
		ok := false
		for _, value := range *res {
			if reflect.DeepEqual(value, stack) {
				ok = true
				break
			}
		}
		if !ok {
			*res = append(*res, stack)
		}
	} else {
		stack = append(stack, nums[pos])
		subsetUniqueBacktrack(res, nums, stack, pos+1)
		stack = utils.Pop(stack)
		subsetUniqueBacktrack(res, nums, stack, pos+1)
	}
}
