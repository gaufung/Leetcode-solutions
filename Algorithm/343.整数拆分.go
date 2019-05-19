/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 *
 * https://leetcode-cn.com/problems/integer-break/description/
 *
 * algorithms
 * Medium (50.30%)
 * Likes:    79
 * Dislikes: 0
 * Total Accepted:    5.4K
 * Total Submissions: 10.6K
 * Testcase Example:  '2'
 *
 * 给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。
 *
 * 示例 1:
 *
 * 输入: 2
 * 输出: 1
 * 解释: 2 = 1 + 1, 1 × 1 = 1。
 *
 * 示例 2:
 *
 * 输入: 10
 * 输出: 36
 * 解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
 *
 * 说明: 你可以假设 n 不小于 2 且不大于 58。
 *
 */
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		val := 1
		for j := 1; j < i; j++ {
			if j*(i-j) > val {
				val = j * (i - j)
			}
			if j*dp[i-j] > val {
				val = j * dp[i-j]
			}
			if dp[j]*(i-j) > val {
				val = dp[j] * (i - j)
			}
			if dp[j]*dp[i-j] > val {
				val = dp[j] * dp[i-j]
			}
		}
		dp[i] = val
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

