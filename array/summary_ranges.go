/*
Given a sorted integer array without duplicates,
return the summary of its ranges.

For example, given [0,1,2,4,5,7], return ["0->2","4->5","7"].
*/

package array

import (
	"bytes"
	"strconv"
)

func summaryRanges(nums []int) []string {
	res := []string{}
	if len(nums) == 1 {
		res = append(res, strconv.Itoa(nums[0]))
		return res
	}
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		for i+1 < len(nums) && nums[i+1]-nums[i] == 1 {
			i += 1
		}
		if nums[i] != num {
			var tempBuffer bytes.Buffer
			tempBuffer.WriteString(strconv.Itoa(num))
			tempBuffer.WriteString("->")
			tempBuffer.WriteString(strconv.Itoa(nums[i]))
			res = append(res, tempBuffer.String())
		} else {
			res = append(res, strconv.Itoa(num))
		}
	}
	return res
}
