/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 *
 * https://leetcode-cn.com/problems/jump-game-ii/description/
 *
 * algorithms
 * Hard (30.10%)
 * Likes:    185
 * Dislikes: 0
 * Total Accepted:    9.4K
 * Total Submissions: 31.1K
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * 给定一个非负整数数组，你最初位于数组的第一个位置。
 *
 * 数组中的每个元素代表你在该位置可以跳跃的最大长度。
 *
 * 你的目标是使用最少的跳跃次数到达数组的最后一个位置。
 *
 * 示例:
 *
 * 输入: [2,3,1,1,4]
 * 输出: 2
 * 解释: 跳到最后一个位置的最小跳跃数是 2。
 * 从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
 *
 *
 * 说明:
 *
 * 假设你总是可以到达数组的最后一个位置。
 *
 */
func jump(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		val := n
		for j := i - 1; j >= 0; j-- {
			if i-j <= nums[j] {
				val = min(val, dp[j])
			}
		}
		dp[i] = val + 1
	}
	return dp[n-1]

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

