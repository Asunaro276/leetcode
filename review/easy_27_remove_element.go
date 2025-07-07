package leetcodereview

import (
	"fmt"
)

/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	writeIndex := 0
	for readIndex := 0; readIndex < len(nums); readIndex++ {
		if nums[readIndex] != val {
			nums[writeIndex] = nums[readIndex]
			writeIndex++
		}
		fmt.Println(nums)
	}
	return writeIndex
}

// func removeElement(nums []int, val int) int {
// 	nums = slices.DeleteFunc(nums, func(n int) bool {
// 		return n == val
// 	})
// 	return len(nums)
// }

// @lc code=end
