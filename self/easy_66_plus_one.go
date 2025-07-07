package leetcode

import (
	"slices"
)

/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

// @lc code=start
func plusOne(digits []int) []int {
	for i := len(digits) - 1; 0 <= i; i-- {
		if digits[i] != 9 {
			digits[i]++
			break
		} else {
			digits[i] = 0
		}
	}
	if digits[0] == 0 {
		digits = slices.Insert(digits, 0, 1)
	}
	return digits
}

// @lc code=end
