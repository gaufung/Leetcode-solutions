import "math"

/*
 * @lc app=leetcode.cn id=507 lang=golang
 *
 * [507] 完美数
 *
 * https://leetcode-cn.com/problems/perfect-number/description/
 *
 * algorithms
 * Easy (35.19%)
 * Likes:    32
 * Dislikes: 0
 * Total Accepted:    6K
 * Total Submissions: 17K
 * Testcase Example:  '28'
 *
 * 对于一个 正整数，如果它和除了它自身以外的所有正因子之和相等，我们称它为“完美数”。
 *
 * 给定一个 正整数 n， 如果他是完美数，返回 True，否则返回 False
 *
 *
 *
 * 示例：
 *
 *
 * 输入: 28
 * 输出: True
 * 解释: 28 = 1 + 2 + 4 + 7 + 14
 *
 *
 *
 *
 * 注意:
 *
 * 输入的数字 n 不会超过 100,000,000. (1e8)
 *
 */
func checkPerfectNumber(num int) bool {
	sum := 1
	if num == 1 {
		return false
	}
	to := int(math.Sqrt(float64(num)))
	for i := 2; i <= to; i++ {
		if num%i == 0 {
			left := num / i
			if i != left {
				sum += i
				sum += left
			} else {
				sum += i
			}
		}

	}
	return sum == num
}

