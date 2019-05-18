/*
 * @lc app=leetcode.cn id=115 lang=golang
 *
 * [115] 不同的子序列
 */
func numDistinct(s string, t string) int {
	m, n := len(t), len(s)
	if m == 0 {
		return 1
	}
	if n == 0 {
		return 0
	}
	if m > n {
		return 0
	}
	if m == 0 {
		return 0
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	if s[0] == t[0] {
		dp[0][0] = 1
	}
	for i := 1; i < n; i++ {
		if t[0] == s[i] {
			dp[0][i] = dp[0][i-1] + 1
		} else {
			dp[0][i] = dp[0][i-1]
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if t[i] == s[j] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

