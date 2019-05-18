/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
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
