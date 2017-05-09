/*
WAP to take one element from each of the array add it to the target sum. Print all those three-element combinations.

A = [1, 2, 3, 3]
B = [2, 3, 3, 4]
C = [1, 2, 2, 2]
target = 7

Result:
[[1, 2, 4], [1, 3, 3], [1, 3, 3], [1, 3, 3], [1, 3, 3], [1, 4, 2], [2, 2, 3], [2, 2, 3], [2, 3, 2], [2, 3, 2], [3, 2, 2], [3, 2, 2]]
*/

package backtrack

import "fmt"

type any []interface{}

var A []int = []int{1, 2, 3, 3}
var B []int = []int{2, 3, 3, 4}
var C []int = []int{1, 2, 2, 2}

func constructCandidates(constructedSofar []int) []int {
	array := A
	if len(constructedSofar) == 1 {
		array = B
	} else if len(constructedSofar) == 2 {
		array = C
	}

	return array
}

func over(constructedSofar []int, target int) (bool, bool) {
	sum := 0
	toStop, reachedTarget := false, false
	for _, value := range constructedSofar {
		sum += value
	}
	if sum >= target || len(constructedSofar) >= 3 {
		toStop = true
		if sum == target && len(constructedSofar) == 3 {
			reachedTarget = true
		}
	}

	return toStop, reachedTarget
}

func sumCombinationsBacktrack(constructedSofar *[]int, target int) {
	toStop, reachedTarget := over(*constructedSofar, target)
	if toStop {
		if reachedTarget {
			fmt.Print(*constructedSofar)
		}
		return
	}
	candidates := constructCandidates(*constructedSofar)
	for _, value := range candidates {
		*constructedSofar = append(*constructedSofar, value)
		sumCombinationsBacktrack(constructedSofar, target)
		*constructedSofar = pop(*constructedSofar)
	}
}

func pop(slice []int) []int {
	slice = slice[:len(slice)-1]
	return slice
}