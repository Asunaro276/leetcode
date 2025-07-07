package leetcode
/*
 * @lc app=leetcode id=59 lang=golang
 *
 * [59] Spiral Matrix II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
  result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}
	top, right, bottom, left := 0, n-1, n-1, 0
	i := 1
	for top <= bottom && left <= right {
		for j := left; j <= right; j++ {
			result[top][j] = i
			i++
		}
		top++

		for j := top; j <= bottom; j++ {
			result[j][right] = i
			i++
		}
		right--

		for j := right; left <= j; j-- {
			result[bottom][j] = i
			i++
		}
		bottom--

		for j := bottom; top <= j; j-- {
			result[j][left] = i
			i++
		}
		left++
	}
	return result
}
// @lc code=end
