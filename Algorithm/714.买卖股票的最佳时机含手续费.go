/*
 * @lc app=leetcode.cn id=714 lang=golang
 *
 * [714] 买卖股票的最佳时机含手续费
 */
func maxProfit(prices []int, fee int) int {
	n := len(prices)
	if n <= 0 {
		return 0
	}
	hold := make([]int, n)
	unhold := make([]int, n)
	hold[0] = -prices[0]
	unhold[0] = 0
	for i := 1; i < n; i++ {
		hold[i] = max(hold[i-1], unhold[i-1]-prices[i])
		unhold[i] = max(unhold[i-1], hold[i-1]+prices[i]-fee)
	}
	return unhold[n-1]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

