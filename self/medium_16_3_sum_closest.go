package leetcode

import (
	"math"
	"slices"
)

/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
  slices.Sort(nums)
  result := nums[0] + nums[1] + nums[len(nums) - 1]
  for point := 0; point < len(nums) - 2; point++ {
    left, right := point + 1, len(nums) - 1
    for left < right {
      sum := nums[point] + nums[left] + nums[right]
      if math.Abs(float64(target - result)) >= math.Abs(float64(target - sum)) {
        result = sum
      }

      if target > sum {
        left++
      } else {
        right--
      }
    }
  }
  return result
}
// @lc code=end
