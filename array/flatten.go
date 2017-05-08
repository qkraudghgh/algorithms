/*
Implement Flatten Arrays.
Given an array that may contain nested arrays,
give a single resultant array.

function flatten(input){
}

Example:

Input: var input = [2, 1, [3, [4, 5], 6], 7, [8]];
flatten(input);
Output: [2, 1, 3, 4, 5, 6, 7, 8]
*/

package array

type any []interface{}

func flatten(list, result any) any {
	for _, value := range list {
		s, ok := value.(int)
		if ok {
			result = append(result, s)
		} else {
			result = flatten(value.(any), result)
		}
	}
	return result
}
