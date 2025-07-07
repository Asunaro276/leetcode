package leetcode

import (
	"slices"
)

/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}

	results := [][]int{}
	slices.Sort(nums)
	for a := 0; a < len(nums)-3; a++ {
		if a > 0 && nums[a-1] == nums[a] {
			continue
		}
		for b := a + 1; b < len(nums)-2; b++ {
			c, d := b+1, len(nums)-1
			if b > a+1 && nums[b-1] == nums[b] {
				continue
			}
			for c < d {
				if b+1 < c && nums[c-1] == nums[c] {
					c++
					continue
				} else if d < len(nums)-1 && nums[d] == nums[d+1] {
					d--
					continue
				}

				sum := nums[a] + nums[b] + nums[c] + nums[d]
				if sum == target {
					results = append(results, []int{nums[a], nums[b], nums[c], nums[d]})
					c++
					d--
				} else if sum < target {
					c++
				} else {
					d--
				}
			}
		}
	}
	return results
}

// @lc code=end
