/*
 * @lc app=leetcode.cn id=485 lang=golang
 *
 * [485] 最大连续1的个数
 *
 * https://leetcode-cn.com/problems/max-consecutive-ones/description/
 *
 * algorithms
 * Easy (53.77%)
 * Likes:    56
 * Dislikes: 0
 * Total Accepted:    14.2K
 * Total Submissions: 26.4K
 * Testcase Example:  '[1,0,1,1,0,1]'
 *
 * 给定一个二进制数组， 计算其中最大连续1的个数。
 *
 * 示例 1:
 *
 *
 * 输入: [1,1,0,1,1,1]
 * 输出: 3
 * 解释: 开头的两位和最后的三位都是连续1，所以最大连续1的个数是 3.
 *
 *
 * 注意：
 *
 *
 * 输入的数组只包含 0 和1。
 * 输入数组的长度是正整数，且不超过 10,000。
 *
 *
 */
func findMaxConsecutiveOnes(nums []int) int {
	result := 0
	first := left(nums, 0)
	second := first
	n := len(nums)
	for second < n {
		if nums[second] == 1 {
			second++
		} else {
			result = max(result, second-first)
			first = left(nums, second)
			second = first
		}
	}
	result = max(result, second-first)
	return result
}

func left(nums []int, from int) int {
	n := len(nums)
	for from < n && nums[from] == 0 {
		from++
	}
	return from
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

