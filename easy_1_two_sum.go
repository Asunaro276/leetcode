package leetcode
/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	filteredNums := []int{}
	for _, n := range nums {
		if n < target {
			filteredNums = append(filteredNums, n)
		}
	}

	for _, n := filteredNums {
		for j := i+1; j < l; j++ {
			if target == nums[i] + nums[j] {
				return []int{i, j}
			}
		} 
	}
	return []int{-1, -1}
}
// @lc code=end
