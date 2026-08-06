package main

import (
	"slices"
	"sort"
)

func sortByEnds(intervals [][]int) {
	slices.SortFunc(intervals, func(a, b []int) int {
		if a[1] < b[1] {
			return -1
		}
		if a[1] > b[1] {
			return 1
		}
		return 0
	})
}

/*
Підраховує і поверає кількість інтервалів `intervals`,
якіх не входять в межі будь якого іншого інтервала.

	removeCoveredIntervals([][]int{
		{1, 4}, // +-----+     ✔
		{3, 6}, //     +-----+ ✔
		{2, 5}, //   +-----+   ✔
		{4, 6}, //       +---+ X Входить в {3, 6}
		        // 1 2 3 4 5 6
	}) -> 3
*/
func removeCoveredIntervals(intervals [][]int) int {
	sortByEnds(intervals)
	seq := make([]bool, len(intervals))

	for i := range intervals {
		start, end := intervals[i][0], intervals[i][1]

		j := sort.Search(len(intervals), func(j int) bool {
			return start < intervals[j][1]
		})

		for ; j < len(intervals) && intervals[j][1] <= end; j++ {
			if start < intervals[j][0] || (start == intervals[j][0] && intervals[j][1] < end) {
				seq[j] = true
			}
		}
	}

	newLen := len(seq)

	for _, isLost := range seq {
		if isLost {
			newLen--
		}
	}
	return newLen
}

func main() {
	println(removeCoveredIntervals([][]int{{1, 4}, {2, 3}}))
}
