package leetcode

import (
	"slices"
)

/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 */

// @lc code=start
func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i > -1 {
		j := len(nums) - 1
		for j > 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	slices.Reverse(nums[i+1:])
}

// @lc code=end
