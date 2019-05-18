/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */
func countBits(num int) []int {
	if num == 0 {
		return []int{0}
	}
	dp := make([]int, num+1)
	dp[0] = 0
	for i := 1; i <= num; i++ {
		dp[i] = dp[(i&(i-1))] + 1
	}
	return dp
}

