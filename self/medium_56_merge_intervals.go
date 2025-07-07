package leetcode

import (
	"slices"
)

/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
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
			result[len(result)-1] = []int{result[len(result)-1][0], max(result[len(result)-1][1], intervals[i][1])}
		} else {
			result = append(result, intervals[i])
		}
	}
	return result
}

// @lc code=end
