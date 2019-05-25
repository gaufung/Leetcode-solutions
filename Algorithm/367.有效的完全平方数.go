/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 *
 * https://leetcode-cn.com/problems/valid-perfect-square/description/
 *
 * algorithms
 * Easy (41.05%)
 * Likes:    49
 * Dislikes: 0
 * Total Accepted:    9.7K
 * Total Submissions: 23.7K
 * Testcase Example:  '16'
 *
 * 给定一个正整数 num，编写一个函数，如果 num 是一个完全平方数，则返回 True，否则返回 False。
 *
 * 说明：不要使用任何内置的库函数，如  sqrt。
 *
 * 示例 1：
 *
 * 输入：16
 * 输出：True
 *
 * 示例 2：
 *
 * 输入：14
 * 输出：False
 *
 *
 */
func isPerfectSquare(num int) bool {
	lo, hi := 1, num+1
	for lo < hi {
		mi := (lo + hi) >> 1
		if mi*mi == num {
			return true
		} else if mi*mi > num {
			hi = mi
		} else {
			lo = mi + 1
		}
	}
	return false
}

