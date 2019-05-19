/*
 * @lc app=leetcode.cn id=357 lang=golang
 *
 * [357] 计算各个位数不同的数字个数
 *
 * https://leetcode-cn.com/problems/count-numbers-with-unique-digits/description/
 *
 * algorithms
 * Medium (44.96%)
 * Likes:    27
 * Dislikes: 0
 * Total Accepted:    2.2K
 * Total Submissions: 4.9K
 * Testcase Example:  '2'
 *
 * 给定一个非负整数 n，计算各位数字都不同的数字 x 的个数，其中 0 ≤ x < 10^n 。
 *
 * 示例:
 *
 * 输入: 2
 * 输出: 91
 * 解释: 答案应为除去 11,22,33,44,55,66,77,88,99 外，在 [0,100) 区间内的所有数字。
 *
 *
 */
func countNumbersWithUniqueDigits(n int) int {
	result := make([]int, 10)
	result[0] = 1
	result[1] = 10
	for i := 2; i <= n; i++ {
		result[i] = result[i-1] + (result[i-1]-result[i-2])*(11-i)
	}
	return result[n]
}

