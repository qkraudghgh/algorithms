/*
Given a collection of numbers that might contain duplicates,
return all possible unique permutations.

For example,
[1,1,2] have the following unique permutations:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
*/

package backtrack

import "fmt"

func permuteUnique(nums []int) [][]int {
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
				if i < len(perm) && perm[i] == num {
					break
				}
				fmt.Println(i, perm[:i], []int{num}, perm[i:], ">>>>", newPerms)
			}
		}
		perms = newPerms
	}
	return perms
}
