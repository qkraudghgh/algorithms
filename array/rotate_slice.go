/*
Rotate an array of n elements to the right by k steps.

For example, with n = 7 and k = 3,
the array [1,2,3,4,5,6,7] is rotated to [5,6,7,1,2,3,4].

Note:
Try to come up as many solutions as you can,
there are at least 3 different ways to solve this problem.
*/

package array

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	reverse(nums, 0, n-k-1)
	reverse(nums, n-k, n-1)
	reverse(nums, 0, n-1)
}

func reverse(slice []int, a int, b int) {
	for a < b {
		slice[a], slice[b] = slice[b], slice[a]
		a += 1
		b -= 1
	}
}
