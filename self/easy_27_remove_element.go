package leetcode

/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	writeIndex := 0
	for readIndex := 0; readIndex < len(nums); readIndex++ {
		if nums[readIndex] != val {
			nums[writeIndex] = nums[readIndex]
			writeIndex++
		}
	}
	return writeIndex
}

// @lc code=end
