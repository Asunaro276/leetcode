package leetcode

/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	n := len(matrix[0])
	m := len(matrix)
	result := make([]int, n*m)
	left, right, top, bottom := 0, n-1, 0, m-1
	j := 0
	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			result[j] = matrix[top][i]
			j++
		}
		top++

		for i := top; i <= bottom; i++ {
			result[j] = matrix[i][right]
			j++
		}
		right--

		if left <= right && top <= bottom {
			for i := right; left <= i; i-- {
				result[j] = matrix[bottom][i]
				j++
			}
			bottom--

			for i := bottom; top <= i; i-- {
				result[j] = matrix[i][left]
				j++
			}
			left++
		}
	}
	return result
}

// @lc code=end
