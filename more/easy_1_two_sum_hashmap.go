package leetcodemore

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum - HashMap Approach
 *
 * 解法: ハッシュマップを使用した最適解
 * 時間計算量: O(n)
 * 空間計算量: O(n)
 *
 * アルゴリズム:
 * 1. ハッシュマップ(map)を作成して、値とインデックスを保存
 * 2. 配列を一度だけ走査
 * 3. 各要素について、complement = target - current_element を計算
 * 4. complementがマップに存在するかチェック
 * 5. 存在すれば解を返す、なければ現在の要素をマップに追加
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	// Step 1: ハッシュマップを初期化
	// key: 配列の要素の値, value: そのインデックス
	valueToIndexMap := make(map[int]int, 50)

	// Step 2: 配列を前から順に走査
	for i := 0; i < len(nums); i++ {
		currentValue := nums[i]

		// Step 3: 現在の値と合わせてtargetになる値（補数）を計算
		requiredValue := target - currentValue

		// Step 4: 補数がマップに存在するかチェック（マップの検索はO(1)でできるので高速）
		if foundIndex, isFound := valueToIndexMap[requiredValue]; isFound {
			// Step 5: 存在する場合、二つのインデックスを返す
			// 注意: foundIndexは必ずiより小さい（先に追加されたため）
			return []int{foundIndex, i}
		}

		// Step 6: 現在の値とインデックスをマップに保存
		valueToIndexMap[currentValue] = i
	}

	// Step 7: 解が見つからない場合（実際には発生しない）
	return []int{}
}
// @lc code=end
