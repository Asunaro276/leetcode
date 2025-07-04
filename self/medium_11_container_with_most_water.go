package leetcode

import "fmt"

/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func maxArea(height []int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("PANIC: %v\n", r)
		}
	}()
	maxArea := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			area := min(height[i], height[j]) * (j - i)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

// @lc code=end
