/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] 判断子序列
 */
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	m, n := len(s), len(t)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
			if dp[i][j] == m {
				return true
			}
		}
	}
	return dp[m][n] == m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}