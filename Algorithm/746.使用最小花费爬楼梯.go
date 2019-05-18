/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] 使用最小花费爬楼梯
 */
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	if n <= 2 {
		return 0
	}
	//dp[i] = min(dp[i-1] + nums[i-1], dp[i-2] + nums[i-2])
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 0
	for i := 2; i < n+1; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
