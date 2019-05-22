/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 *
 * https://leetcode-cn.com/problems/minimum-size-subarray-sum/description/
 *
 * algorithms
 * Medium (38.50%)
 * Likes:    97
 * Dislikes: 0
 * Total Accepted:    9.7K
 * Total Submissions: 25.2K
 * Testcase Example:  '7\n[2,3,1,2,4,3]'
 *
 * 给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组。如果不存在符合条件的连续子数组，返回 0。
 *
 * 示例:
 *
 * 输入: s = 7, nums = [2,3,1,2,4,3]
 * 输出: 2
 * 解释: 子数组 [4,3] 是该条件下的长度最小的连续子数组。
 *
 *
 * 进阶:
 *
 * 如果你已经完成了O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。
 *
 */
func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}
	partialSum := make([]int, n)
	partialSum[0] = nums[0]
	for i := 1; i < n; i++ {
		partialSum[i] = partialSum[i-1] + nums[i]
	}
	start := findLargeOrEqual(partialSum, s)
	if start == n {
		return 0
	}
	res := n
	for start < n {
		left := partialSum[start] - s
		from := findLessOrEqual(partialSum, left, 0, start)
		if start-from < res {
			res = start - from
		}
		start++
	}
	return res

}

func findLessOrEqual(nums []int, target int, lo, hi int) int {
	for lo < hi {
		mi := (hi + lo) >> 1
		if nums[mi] > target {
			hi = mi
		} else {
			lo = mi + 1
		}
	}
	return lo - 1
}

func findLargeOrEqual(nums []int, target int) int {
	lo, hi := 0, len(nums)
	for lo < hi {
		mi := (hi + lo) >> 1
		if nums[mi] < target {
			lo = mi + 1
		} else {
			hi = mi
		}
	}
	return lo
}

