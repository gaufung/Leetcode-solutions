/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 *
 * https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/description/
 *
 * algorithms
 * Medium (46.80%)
 * Likes:    83
 * Dislikes: 0
 * Total Accepted:    3.2K
 * Total Submissions: 6.8K
 * Testcase Example:  '[1,2,3,0,2]'
 *
 * 给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​
 *
 * 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
 *
 *
 * 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
 * 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
 *
 *
 * 示例:
 *
 * 输入: [1,2,3,0,2]
 * 输出: 3
 * 解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
 *
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

