/*
 * @lc app=leetcode.cn id=201 lang=golang
 *
 * [201] 数字范围按位与
 *
 * https://leetcode-cn.com/problems/bitwise-and-of-numbers-range/description/
 *
 * algorithms
 * Medium (39.43%)
 * Likes:    34
 * Dislikes: 0
 * Total Accepted:    2.5K
 * Total Submissions: 6.3K
 * Testcase Example:  '5\n7'
 *
 * 给定范围 [m, n]，其中 0 <= m <= n <= 2147483647，返回此范围内所有数字的按位与（包含 m, n 两端点）。
 *
 * 示例 1:
 *
 * 输入: [5,7]
 * 输出: 4
 *
 * 示例 2:
 *
 * 输入: [0,1]
 * 输出: 0
 *
 */
func rangeBitwiseAnd(m int, n int) int {
	rightShift := uint(0)
	res := 0
	for rightShift < 32 {
		mm := m >> rightShift
		nn := n >> rightShift
		bit := 1
		if nn-mm >= 1 {
			bit = 0
		} else {
			if nn%2 == 0 {
				bit = 0
			}
		}
		res = res | (bit << rightShift)
		rightShift++
	}
	return res
}

