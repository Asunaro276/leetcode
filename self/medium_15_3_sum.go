package leetcode
/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	results := make([][]int, 0)
	for i := 0; i < len(nums) - 1; i++ {
		hashMap := make(map[int]int)
		for j := i + 1; j < len(nums); j ++ {
			currentValue := nums[j]
			complement := - (nums[i] + nums[j])
			if foundIndex, isFound := hashMap[complement]; isFound {
				results = append(results, []int{nums[i], nums[j], nums[foundIndex]})
			}
			hashMap[currentValue] = j
		}
	}
	return results
}
// @lc code=end
