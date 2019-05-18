/*
 * @lc app=leetcode.cn id=375 lang=golang
 *
 * [375] 猜数字大小 II
 */
func getMoneyAmount(n int) int {
	dp := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		dp[i] = make([]int, n+2)
	}
	maxMoney := n * (n - 1) / 2
	for i := n; i >= 1; i-- {
		for j := i; j <= n; j++ {
			if i == j {
				dp[i][j] = 0
			} else {
				dp[i][j] = maxMoney
				for k := i; k <= j; k++ {
					dp[i][j] = min(dp[i][j], max(dp[i][k-1], dp[k+1][j])+k)
				}
			}
		}
	}
	return dp[1][n]

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

