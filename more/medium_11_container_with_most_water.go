package leetcodemore

/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxWater := 0

	for left < right {
		// 現在の水の量を計算
		currentArea := min(height[left], height[right]) * (right - left)

		// 最大値を更新
		if currentArea > maxWater {
			maxWater = currentArea
		}

		// 高さが低い方のポインタを内側に移動
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxWater
}
// @lc code=end
