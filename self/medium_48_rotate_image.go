package leetcode

/*
 * @lc app=leetcode id=48 lang=golang
 *
 * [48] Rotate Image
 */

// @lc code=start
func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-i-1] = matrix[n-i-1], matrix[i]
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			matrix[n-i-1][j], matrix[j][n-i-1] = matrix[j][n-i-1], matrix[n-i-1][j]
		}
	}
}

// @lc code=end
