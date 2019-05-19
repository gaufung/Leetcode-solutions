/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长上升子序列
 *
 * https://leetcode-cn.com/problems/longest-increasing-subsequence/description/
 *
 * algorithms
 * Medium (41.26%)
 * Likes:    180
 * Dislikes: 0
 * Total Accepted:    12.8K
 * Total Submissions: 31K
 * Testcase Example:  '[10,9,2,5,3,7,101,18]'
 *
 * 给定一个无序的整数数组，找到其中最长上升子序列的长度。
 *
 * 示例:
 *
 * 输入: [10,9,2,5,3,7,101,18]
 * 输出: 4
 * 解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
 *
 * 说明:
 *
 *
 * 可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
 * 你算法的时间复杂度应该为 O(n^2) 。
 *
 *
 * 进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?
 *
 */
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}
	lowest := make([]int, n)
	lowest[0] = nums[0]
	length := 1
	for i := 1; i < n; i++ {
		if nums[i] > lowest[length-1] {
			lowest[length] = nums[i]
			length++
		} else {
			j := binarySearch(lowest, nums[i], length) + 1
			lowest[j] = nums[i]
		}
	}
	return length
}

func binarySearch(nums []int, target int, n int) int {
	lo, hi := 0, n
	for lo < hi {
		mi := (lo + hi) >> 1
		if nums[mi] < target {
			lo = mi + 1
		} else {
			hi = mi
		}
	}
	return hi - 1
}

