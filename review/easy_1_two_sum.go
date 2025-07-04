package leetcodereview
/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
  hashMap := make(map[int]int)
	for i, v := range nums {
		complement := target - v
		if foundIndex, isFound := hashMap[complement]; isFound {
			return []int{foundIndex, i}
		}
		hashMap[v] = i
	}
	return []int{}
}
// @lc code=end
