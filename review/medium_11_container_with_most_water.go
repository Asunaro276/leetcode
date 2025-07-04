package leetcodereview

/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	result := 0
	for left < right {
		if height[left] > height[right] {
			area := height[right] * (right - left)
			if area > result {
				result = area
			} else {
				right--
			}
		} else {
			area := height[left] * (right - left)
			if area > result {
				result = area
			} else {
				left++
			}
		}
	}
	return result
}

// @lc code=end
