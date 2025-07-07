package leetcode

import "strings"

/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	table := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	result := 0
	for i, c := range s {
		minusCond := i > 0 && ((strings.Contains("I", string(s[i-1])) && strings.Contains("VX", string(s[i]))) || (strings.Contains("X", string(s[i-1])) && strings.Contains("LC", string(s[i]))) || (strings.Contains("C", string(s[i-1])) && strings.Contains("DM", string(s[i]))))
		if minusCond {
			result -= 2 * table[string(s[i-1])]	
		}
		result += table[string(c)]	
	}
	return result
}
// @lc code=end
