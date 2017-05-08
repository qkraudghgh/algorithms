/*
find missing ranges between low and high in the given array.
ex) [3, 5] lo=1 hi=10 => answer: [1->2, 4, 6->10]
*/

package array

import (
	"bytes"
	"strconv"
)

func missingRanges(nums []int, low, high int) []string {
	res := []string{}
	start := low
	for _, num := range nums {
		if num < start {
			continue
		}
		if num == start {
			start += 1
			continue
		}
		res = append(res, getRange(start, num-1))
		start = num + 1
	}
	if start <= high {
		res = append(res, getRange(start, high))
	}
	return res
}

func getRange(n1 int, n2 int) string {
	if n1 == n2 {
		return strconv.Itoa(n1)
	} else {
		var tempBuffer bytes.Buffer
		tempBuffer.WriteString(strconv.Itoa(n1))
		tempBuffer.WriteString("->")
		tempBuffer.WriteString(strconv.Itoa(n2))
		return tempBuffer.String()
	}
}
