/*
Given a string that contains only digits 0-9 and a target value,
return all possibilities to add binary operators (not unary) +, -, or *
between the digits so they prevuate to the target value.

Examples:
"123", 6 -> ["1+2+3", "1*2*3"]
"232", 8 -> ["2*3+2", "2+3*2"]
"105", 5 -> ["1*0+5","10-5"]
"00", 0 -> ["0+0", "0-0", "0*0"]
"3456237490", 9191 -> []
*/

package backtrack

import (
	"strconv"
)

func addOperator(num string, target int) []string {
	res := []string{}
	if len(num) == 0 {
		return res
	}
	helper(&res, "", num, target, 0, 0, 0)
	return res
}

func helper(res *[]string, path string, num string, target int, pos int, prev int, multed int) {
	if pos == len(num) {
		if target == prev {
			*res = append(*res, path)
		}
		return
	}
	for i := pos; i < len(num); i++ {
		if i != pos && num[pos] == '0' {
			break
		}
		cur, _ := strconv.Atoi(num[pos : i+1])
		if pos == 0 {
			helper(res, path+strconv.Itoa(cur), num, target, i+1, cur, cur)
		} else {
			helper(res, path+"+"+strconv.Itoa(cur), num, target, i+1, prev+cur, cur)
			helper(res, path+"-"+strconv.Itoa(cur), num, target, i+1, prev-cur, -cur)
			helper(res, path+"*"+strconv.Itoa(cur), num, target, i+1, prev-multed+multed*cur, multed*cur)
		}
	}
}
