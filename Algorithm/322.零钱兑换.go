/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 *
 * https://leetcode-cn.com/problems/coin-change/description/
 *
 * algorithms
 * Medium (31.92%)
 * Likes:    132
 * Dislikes: 0
 * Total Accepted:    10.2K
 * Total Submissions: 31.9K
 * Testcase Example:  '[1,2,5]\n11'
 *
 * 给定不同面额的硬币 coins 和一个总金额
 * amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
 *
 * 示例 1:
 *
 * 输入: coins = [1, 2, 5], amount = 11
 * 输出: 3
 * 解释: 11 = 5 + 5 + 1
 *
 * 示例 2:
 *
 * 输入: coins = [2], amount = 3
 * 输出: -1
 *
 * 说明:
 * 你可以认为每种硬币的数量是无限的。
 *
 */
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		min := -1
		found := false
		for _, c := range coins {
			if i >= c && dp[i-c] != -1 {
				if !found {
					min = dp[i-c] + 1
					found = true
				} else {
					if dp[i-c]+1 < min {
						min = dp[i-c] + 1
					}
				}
			}
		}
		dp[i] = min
	}
	return dp[amount]
}

