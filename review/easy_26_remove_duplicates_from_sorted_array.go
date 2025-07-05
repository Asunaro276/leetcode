package leetcodereview

/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast-1] < nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}

// @lc code=end
