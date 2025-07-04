package leetcodemore

import "sort"

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	results := make([][]int, 0)
	n := len(nums)

	// 配列をソートする
	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		// 重複をスキップ
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				results = append(results, []int{nums[i], nums[left], nums[right]})

				// 重複をスキップ
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return results
}
// @lc code=end
