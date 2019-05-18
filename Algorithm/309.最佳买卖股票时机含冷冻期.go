/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */
func maxProfit(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	if n == 2 {
		return max(prices[1]-prices[0], 0)
	}
	hold, unhold := make([]int, n), make([]int, n)
	hold[0] = -prices[0]
	unhold[0] = 0
	hold[1] = max(-prices[0], -prices[1])
	unhold[1] = max(0, prices[1]-prices[0])
	for i := 2; i < n; i++ {
		hold[i] = max(hold[i-1], unhold[i-2]-prices[i])
		unhold[i] = max(unhold[i-1], hold[i-1]+prices[i])
	}
	return unhold[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

