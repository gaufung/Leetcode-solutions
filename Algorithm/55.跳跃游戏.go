/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 *
 * https://leetcode-cn.com/problems/jump-game/description/
 *
 * algorithms
 * Medium (35.03%)
 * Likes:    180
 * Dislikes: 0
 * Total Accepted:    17.4K
 * Total Submissions: 49.7K
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * 给定一个非负整数数组，你最初位于数组的第一个位置。
 *
 * 数组中的每个元素代表你在该位置可以跳跃的最大长度。
 *
 * 判断你是否能够到达最后一个位置。
 *
 * 示例 1:
 *
 * 输入: [2,3,1,1,4]
 * 输出: true
 * 解释: 从位置 0 到 1 跳 1 步, 然后跳 3 步到达最后一个位置。
 *
 *
 * 示例 2:
 *
 * 输入: [3,2,1,0,4]
 * 输出: false
 * 解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
 *
 *
 */
func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	dp := make([]int, len(nums))
	dp[0] = 0 + nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] < i {
			dp[i] = dp[i-1]
		} else if dp[i-1] == i {
			dp[i] = i + nums[i]
		} else {
			dp[i] = max(dp[i-1], i+nums[i])
		}
	}
	return dp[len(nums)-1] >= (len(nums) - 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

