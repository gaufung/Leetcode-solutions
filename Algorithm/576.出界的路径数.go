/*
 * @lc app=leetcode.cn id=576 lang=golang
 *
 * [576] 出界的路径数
 */
func findPaths(m int, n int, N int, i int, j int) int {
	dp := make([][][]int, N+1)
	for k := 0; k < N+1; k++ {
		dp[k] = buildMap(m, n)
	}
	dp[0][i][j] = 1
	cnt := 0
	for k := 0; k < N; k++ {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if dp[k][i][j] != 0 {
					if i-1 < 0 {
						cnt += dp[k][i][j]
					} else {
						dp[k+1][i-1][j] += dp[k][i][j]
					}
					if i+1 >= m {
						cnt += dp[k][i][j]
					} else {
						dp[k+1][i+1][j] += dp[k][i][j]
					}
					if j-1 < 0 {
						cnt += dp[k][i][j]
					} else {
						dp[k+1][i][j-1] += dp[k][i][j]
					}
					if j+1 >= n {
						cnt += dp[k][i][j]
					} else {
						dp[k+1][i][j+1] += dp[k][i][j]
					}
				}
			}
		}
	}
	threshold := 1000000000 + 7
	if cnt > threshold {
		return cnt % threshold
	}
	return cnt

}

func buildMap(m, n int) [][]int {
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
	}
	return result
}
