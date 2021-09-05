package main

import "fmt"

/*
 * 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
 */
func rotate(nums []int, k int) {
	l := len(nums)
	fmt.Println(l)
	if k-1 < l {
		copy(nums, append(nums[l-k:], nums[:l-k]...))
	}
}

/*
 * 给你一个按 非递减顺序 排序的整数数组 nums，
 * 返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
 */
func sortedSquares(nums []int) []int {
	return []int{0}
}

/*
 * 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
 * 如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
 *
 * 请必须使用时间复杂度为 O(log n) 的算法。
 */
func searchInsert(nums []int, target int) int {
	start, end := 0, len(nums)-1

	for end-start >= 0 {
		m := (start + end) / 2
		v := nums[m]
		if v > target {
			end = m - 1
		} else if v < target {
			start = m + 1
		} else {
			return m
		}
	}

	return start
}

/*
 * 你是产品经理，目前正在带领一个团队开发新的产品。
 * 不幸的是，你的产品的最新版本没有通过质量检测。
 * 由于每个版本都是基于之前的版本开发的，所以错误的版本之后的所有版本都是错的。
 *
 * 假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。
 *
 * 你可以通过调用 bool isBadVersion(version) 接口来判断版本号 version 是否在单元测试中出错。
 * 实现一个函数来查找第一个错误的版本。你应该尽量减少对调用 API 的次数。
 */
func isBadVersion(version int) bool {
	return version >= 4
}
func firstBadVersion(n int) int {
	start, end := 0, n

	for end-start >= 0 {
		m := (start + end) / 2
		if isBadVersion(m) {
			end = m - 1
		} else {
			start = m + 1
		}
	}

	return start
}

/*
 * 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target，
 * 写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
 */
func search(nums []int, target int) int {
	start, end := 0, len(nums)-1

	for end-start >= 0 {
		m := (start + end) / 2
		v := nums[m]
		if v > target {
			end = m - 1
		} else if v < target {
			start = m + 1
		} else {
			return m
		}
	}

	return -1
}

/*
 * 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
 */
func lengthOfLongestSubstring(s string) int {
	var start, end, maxStart, maxEnd int
	m := make(map[rune]int)

	for i, k := range s {
		if v, ok := m[k]; ok && v >= start {
			start = v + 1
		} else {
			end = i
		}
		m[k] = i
		if end-start >= maxEnd-maxStart {
			maxStart = start
			maxEnd = end + 1
		}
	}

	return maxEnd - maxStart
}

func main() {
	// fmt.Println(lengthOfLongestSubstring("abcabcdaa"))
	// fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))
	// fmt.Println(firstBadVersion(5))
	// fmt.Println(searchInsert([]int{1, 3, 5, 6}, 1))
	rotate([]int{1, 2}, 3)
}
