/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m <= 0 {
		return n
	}
	if n <= 0 {
		return m
	}
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < n+1; i++ {
		dp[0][i] = i
	}
	for i := 0; i < m+1; i++ {
		dp[i][0] = i
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))
			}
		}
	}
	return dp[m][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

