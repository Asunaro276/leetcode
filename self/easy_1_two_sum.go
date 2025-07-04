package leetcode

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	l := len(nums)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if target == nums[i]+nums[j] {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}

// @lc code=end
