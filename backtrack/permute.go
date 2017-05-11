/*
Given a collection of distinct numbers, return all possible permutations.

For example,
[1,2,3] have the following permutations:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
*/

package backtrack

import "fmt"

func permute(nums []int) [][]int {
	perms := [][]int{}
	perms = append(perms, []int{})
	for _, num := range nums {
		newPerms := [][]int{}
		for _, perm := range perms {
			for i := 0; i <= len(perm); i++ {
				tempSlice := []int{}
				if len(perm[:i]) > 0 {
					for _, value := range perm[:i] {
						tempSlice = append(tempSlice, value)
					}
				}
				tempSlice = append(tempSlice, num)
				if len(perm[i:]) > 0 {
					for _, value := range perm[i:] {
						tempSlice = append(tempSlice, value)
					}
				}
				newPerms = append(newPerms, tempSlice)
				fmt.Println(i, perm[:i], []int{num}, perm[i:], ">>>>", newPerms)
			}
		}
		perms = newPerms
	}
	return perms
}
