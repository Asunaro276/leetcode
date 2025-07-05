package leetcode

/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	if target < nums[left] {
		return 0
	}
	if nums[right] < target {
		return right + 1
	}
	mid := (left + right) / 2
	for left <= right {
		if target < nums[left] {
			if left-mid == 1 {
				return left
			}
			mid, left = left, mid
		}
		if nums[right] < target {
			if mid-right == 1 {
				return mid
			}
			right, mid = mid, right
		}

		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// @lc code=end
