// Given a collection of intervals, merge all overlapping intervals.
//
// For example,
// Given [1,3],[2,6],[8,10],[15,18]
// return [1,6],[8,10],[15,18].

package array

import (
	"sort"
)

// Definition for an interval.
type interval struct {
	start int
	end   int
}

type intervals []interval

func (list intervals) Len() int               { return len(list) }
func (list intervals) Less(i int, j int) bool { return list[i].start < list[j].start }
func (list intervals) Swap(i int, j int)      { list[i], list[j] = list[j], list[i] }

func merge(list intervals) intervals {
	out := []interval{}
	sort.Sort(list)

	for _, interval := range list {
		if len(out) > 0 && interval.start <= out[len(out)-1].end {
			if interval.end > out[len(out)-1].end {
				out[len(out)-1].end = interval.end
			}
		} else {
			out = append(out, interval)
		}
	}

	return out
}
