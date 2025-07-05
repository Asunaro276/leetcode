package leetcodereview

import "slices"

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	result := [][]int{}
	for point := 0; point < len(nums) - 2; point++ {
		left, right := point + 1, len(nums) - 1
		if point > 0 && nums[point-1] == nums[point] {
			continue
		}
		for left < right {
			value := nums[point] + nums[left] + nums[right]
			if value == 0 {
				result = append(result, []int{nums[point], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
					continue
				}
				for left < right && nums[right+1] == nums[right] {
					right--
					continue
				}
			} else if value < 0 {
				left ++		
			} else {
				right--
			}
		}
	}
	return result
}
// @lc code=end
