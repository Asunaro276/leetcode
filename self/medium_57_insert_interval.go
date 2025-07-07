package leetcode

import (
	"slices"
)

/*
 * @lc app=leetcode id=57 lang=golang
 *
 * [57] Insert Interval
 */

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	slices.SortFunc(intervals, func(a []int, b []int) int {
		if a[0] < b[0] {
			return -1
		} else if a[0] > b[0] {
			return 1
		}
		return 0
	})
	result := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		if result[len(result)-1][1] >= intervals[i][0] {
			result[len(result)-1] = []int{result[len(result)-1][0], max(intervals[i][1], result[len(result)-1][1])}
		} else {
			result = append(result, intervals[i])
		}
	}
	return result
}

// @lc code=end
