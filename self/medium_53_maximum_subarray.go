package leetcode

/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	const maxNum = 10000
	left := 0
	sum := -maxNum
	maxSum := -maxNum
	for right := 0; right < len(nums); right++ {
		if sum < 0 && nums[left] < nums[left+1] {
			left++
			sum -= nums[left]
		}

		sum += nums[right]
		if sum <= 0 {
			left = right
			sum = nums[left]
		}
		maxSum = max(maxSum, sum)
	}
	return maxSum
}

// @lc code=end
