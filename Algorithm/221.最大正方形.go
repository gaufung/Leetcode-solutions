/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
 */
func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	if m <= 0 {
		return 0
	}
	n := len(matrix[0])
	if n <= 0 {
		return 0
	}
	dp := make([][]int, m)
	maxLength := 0
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			maxLength = max(maxLength, dp[i][0])
		}
	}
	for i := 0; i < n; i++ {
		if matrix[0][i] == '1' {
			dp[0][i] = 1
			maxLength = max(maxLength, dp[0][i])
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '0' {
				dp[i][j] = 0
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))
			}
			maxLength = max(maxLength, dp[i][j])
		}
	}
	return maxLength * maxLength
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}