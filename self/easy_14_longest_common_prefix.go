package leetcode

/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	// エッジケース：空の配列
	if len(strs) == 0 {
		return ""
	}

	// エッジケース：1つの文字列のみ
	if len(strs) == 1 {
		return strs[0]
	}

	// 最初の文字列を基準とする
	prefix := strs[0]

	// 他の全ての文字列と比較
	for i := 1; i < len(strs); i++ {
		// 共通接頭辞の長さを見つける
		j := 0
		for j < len(prefix) && j < len(strs[i]) && prefix[j] == strs[i][j] {
			j++
		}
		// 共通接頭辞を更新
		prefix = prefix[:j]

		// 空になった場合は早期終了
		if prefix == "" {
			return ""
		}
	}

	return prefix
}

// @lc code=end
