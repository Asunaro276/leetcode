package leetcode

/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
  findCombination(candidates, target, 0, []int{}, &result)
	return result
}

func findCombination(candidates []int, target int, start int, current []int, result *[][]int) {
	if target == 0 {
		*result = append(*result, append([]int{}, current...))
		return
	} else if target < 0 {
		return
	}
	for i := start; i < len(candidates); i++ {
		current = append(current, candidates[i])
		findCombination(candidates, target-candidates[i], i, current, result)
		current = current[:len(current)-1]
	}
}
// @lc code=end
