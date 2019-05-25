/*
 * @lc app=leetcode.cn id=231 lang=golang
 *
 * [231] 2的幂
 *
 * https://leetcode-cn.com/problems/power-of-two/description/
 *
 * algorithms
 * Easy (45.86%)
 * Likes:    111
 * Dislikes: 0
 * Total Accepted:    20.1K
 * Total Submissions: 43.8K
 * Testcase Example:  '1'
 *
 * 给定一个整数，编写一个函数来判断它是否是 2 的幂次方。
 *
 * 示例 1:
 *
 * 输入: 1
 * 输出: true
 * 解释: 2^0 = 1
 *
 * 示例 2:
 *
 * 输入: 16
 * 输出: true
 * 解释: 2^4 = 16
 *
 * 示例 3:
 *
 * 输入: 218
 * 输出: false
 *
 */
func isPowerOfTwo(n int) bool {
	cnt := 0
	mask := 0x01
	for n > 0 {
		if n&mask == 1 {
			cnt++
		}
		n = n >> 1
	}
	return cnt == 1
}

