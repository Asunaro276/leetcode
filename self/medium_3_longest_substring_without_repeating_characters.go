package leetcode

/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	left := 0
	window := make(map[byte]int)
	maxLength := 0
	for right := 0; right < len(s); right++ {
		if _, isFound := window[s[right]]; isFound {
			left = max(left, window[s[right]]+1)
		}
		window[s[right]] = right

		maxLength = max(maxLength, right-left+1)
	}
	return maxLength
}

// @lc code=end
