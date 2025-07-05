package leetcodereview

import (
	"slices"
)

/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 */

// @lc code=start
func nextPermutation(nums []int)  {
	i := len(nums) - 1
	for i > 0 && nums[i-1] >= nums[i] {
		i--
	}

	if i == 0 {
		slices.Reverse(nums)
		return
	}
	
	if i == len(nums) - 1 {
		slices.Reverse(nums[len(nums)-2:])
	}
	
	for j := len(nums) - 1; i <= j; j-- {
		if nums[i-1] < nums[j] {
			nums[i-1], nums[j] = nums[j], nums[i-1]
			break
		}
	}
	slices.Reverse(nums[i:])
}
// @lc code=end
