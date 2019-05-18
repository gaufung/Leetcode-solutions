/*
 * @lc app=leetcode.cn id=718 lang=golang
 *
 * [718] 最长重复子数组
 */
func findLength(A []int, B []int) int {
	m, n := len(A), len(B)
	if m <= 0 || n <= 0 {
		return 0
	}
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	val := 0
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > val {
					val = dp[i][j]
				}
			} else {
				dp[i][j] = 0
			}
		}
	}
	return val
}

