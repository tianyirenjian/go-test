package main

import (
	"fmt"
	// "sort"
)

func main() {
	areas := [][]int{{14, 19}, {5, 12}, {2, 3}, {5, 10}, {1, 7}, {6, 8}, {7, 12}, {13, 19}, {14, 17}}
	// areas := [][]int{{1,3},{2,6},{8,10},{15,18}}
	fmt.Println(merge(areas))
}

type twoDSlice [][]int

func (t twoDSlice) Len() int {
	return len(t)
}

func (t twoDSlice) Swap(i, j int)  {
	t[i], t[j] = t[j], t[i]
}

func (t twoDSlice) Less(i, j int) bool {
	if t[i][0] == t[j][0] {
		return t[i][1] < t[j][1]
	}
	return t[i][0] < t[j][0]
}

func merge(intervals [][]int) [][]int {
	// sort.Sort(twoDSlice(intervals))
	sort(intervals)
	for i := 0; i < len(intervals); i ++ {
		j := i + 1
		if j == len(intervals) {
			break
		}
		if intervals[i][1] >= intervals[j][0] {
			if intervals[i][1] < intervals[j][1] {
				intervals[i][1] = intervals[j][1]
			}
			intervals = append(intervals[:j], intervals[j+1:]...)
			i --
		}
	}
	return intervals
}

func sort(intervals [][]int) {
	length := len(intervals)
	for i := 0; i < length; i ++ {
		for j := i + 1; j < length; j ++ {
			if more(intervals, i, j) {
				intervals[i], intervals[j] = intervals[j], intervals[i]
			}
		}
	}
}
func more(intervals [][]int, i, j int) bool {
	if intervals[i][0] == intervals[j][0] {
		return intervals[i][1] > intervals[j][1]
	}
	return intervals[i][0] > intervals[j][0]
}
