package leetcode

/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	first := findFirst(nums, target)
	if first == -1 {
		return []int{-1, -1}
	}
	last := findLast(nums, target)
	return []int{first , last}
}

func findFirst(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	result := -1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			result = mid
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}

func findLast(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	result := -1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			result = mid
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}
// @lc code=end
