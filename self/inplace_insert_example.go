package leetcode

import "fmt"

// Example 1: Insert at specific index (expanding slice)
func insertAtIndex(nums []int, index int, val int) []int {
	// スライスを1つ拡張
	nums = append(nums, 0)

	// 挿入位置以降の要素を右にシフト
	copy(nums[index+1:], nums[index:])

	// 指定位置に値を挿入
	nums[index] = val

	return nums
}

// Example 2: Insert multiple values at once
func insertMultipleAtIndex(nums []int, index int, values []int) []int {
	// 必要な分だけスライスを拡張
	nums = append(nums, make([]int, len(values))...)

	// 挿入位置以降の要素を右にシフト
	copy(nums[index+len(values):], nums[index:])

	// 値を挿入
	copy(nums[index:], values)

	return nums
}

// Example 3: Insert in sorted order (maintaining sorted array)
func insertInSortedOrder(nums []int, val int) []int {
	// 挿入位置を見つける
	insertIndex := 0
	for insertIndex < len(nums) && nums[insertIndex] < val {
		insertIndex++
	}

	// 挿入位置に値を挿入
	return insertAtIndex(nums, insertIndex, val)
}

// Example 4: In-place insertion with fixed size array (no expansion)
// 既存の要素を上書きしながら挿入
func insertWithOverwrite(nums []int, index int, val int) {
	if index < len(nums) {
		// 挿入位置以降の要素を右にシフト（最後の要素は失われる）
		copy(nums[index+1:], nums[index:len(nums)-1])
		nums[index] = val
	}
}

// Example 5: Two-pointer technique for conditional insertion
// 条件に基づいて要素を挿入/保持
func insertConditionally(nums []int, shouldInsert func(int) bool, newVal int) []int {
	writeIndex := 0

	for readIndex := 0; readIndex < len(nums); readIndex++ {
		if shouldInsert(nums[readIndex]) {
			// 新しい値を挿入
			nums = append(nums, 0) // スライスを拡張
			copy(nums[writeIndex+1:], nums[writeIndex:])
			nums[writeIndex] = newVal
			writeIndex++
		}

		// 元の値を保持
		nums[writeIndex] = nums[readIndex]
		writeIndex++
	}

	return nums[:writeIndex]
}

// デモンストレーション関数
func demonstrateInsertions() {
	fmt.Println("=== In-place Insertion Examples ===")

	// Example 1
	nums1 := []int{1, 2, 4, 5}
	fmt.Printf("Before: %v\n", nums1)
	nums1 = insertAtIndex(nums1, 2, 3)
	fmt.Printf("After inserting 3 at index 2: %v\n", nums1)

	// Example 2
	nums2 := []int{1, 2, 6, 7}
	fmt.Printf("Before: %v\n", nums2)
	nums2 = insertMultipleAtIndex(nums2, 2, []int{3, 4, 5})
	fmt.Printf("After inserting [3,4,5] at index 2: %v\n", nums2)

	// Example 3
	nums3 := []int{1, 3, 5, 7}
	fmt.Printf("Before: %v\n", nums3)
	nums3 = insertInSortedOrder(nums3, 4)
	fmt.Printf("After inserting 4 in sorted order: %v\n", nums3)
}
