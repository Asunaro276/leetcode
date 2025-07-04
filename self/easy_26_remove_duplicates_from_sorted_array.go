package leetcode

/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
// func removeDuplicates(nums []int) int {
//   set := make(map[int]int)
// 	for i, v := range nums {
// 		if _, isFound := set[v]; !isFound {
// 			set[v] = i
// 		} else {
// 			nums[i] = 30001
// 		}
// 	}
// 	slices.Sort(nums)
// 	return len(set)
// }
// @lc code=end
