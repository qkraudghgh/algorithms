/*
There is a parking lot with only one empty spot. Given the initial state
of the parking lot and the final state. Each step we are only allowed to
move a car
out of its place and move it into the empty spot.
The goal is to find out the least movement needed to rearrange
the parking lot from the initial state to the final state.

Say the initial state is an array:

[1,2,3,0,4],
where 1,2,3,4 are different cars, and 0 is the empty spot.

And the final state is

[0,3,2,1,4].
We can swap 1 with 0 in the initial array to get [0,2,3,1,4] and so on.
Each step swap with 0 only.
Edited by cyberking-saga
*/

package array

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

func garage(beg []int, end []int) int {
	i := 0
	moves := 0
	for !reflect.DeepEqual(beg, end) {
		if beg[i] != 0 && beg[i] != end[i] {
			car := beg[i]
			empty, err := findElement(beg, 0)
			if err != nil {
				log.Fatal(err)
			}
			finalPos, err := findElement(end, beg[i])
			if err != nil {
				log.Fatal(err)
			}
			if empty != finalPos {
				beg[finalPos], beg[empty] = beg[empty], beg[finalPos]
				fmt.Println(beg)
				empty, err = findElement(beg, 0)
				if err != nil {
					log.Fatal(err)
				}
				carLotate, err := findElement(beg, car)
				if err != nil {
					log.Fatal(err)
				}
				beg[carLotate], beg[empty] = beg[empty], beg[carLotate]
				fmt.Println(beg)
				moves += 2
			} else {
				carLotate, err := findElement(beg, car)
				if err != nil {
					log.Fatal(err)
				}
				beg[carLotate], beg[empty] = beg[empty], beg[carLotate]
				fmt.Println(beg)
				moves += 1
			}
		}
		i += 1
		if i == len(beg) {
			i = 0
		}
	}
	return moves
}

func findElement(list []int, element int) (int, error) {
	for index, value := range list {
		if value == element {
			return index, nil
		}
	}
	return -1, errors.New("Did not find Element.")
}
