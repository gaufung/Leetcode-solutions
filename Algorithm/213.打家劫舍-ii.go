/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 *
 * https://leetcode-cn.com/problems/house-robber-ii/description/
 *
 * algorithms
 * Medium (32.81%)
 * Likes:    95
 * Dislikes: 0
 * Total Accepted:    5.5K
 * Total Submissions: 16.8K
 * Testcase Example:  '[2,3,2]'
 *
 *
 * 你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都围成一圈，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
 *
 * 给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。
 *
 * 示例 1:
 *
 * 输入: [2,3,2]
 * 输出: 3
 * 解释: 你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
 *
 *
 * 示例 2:
 *
 * 输入: [1,2,3,1]
 * 输出: 4
 * 解释: 你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
 * 偷窃到的最高金额 = 1 + 3 = 4 。
 *
 */
func rob(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	withFirst := nums[1:]
	withLast := nums[:n-1]
	withFirstLast := nums[1 : n-1]
	return max(max(robbery(withFirst), robbery(withLast)),
		robbery(withFirstLast))
}

func robbery(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}
	robbery := make([]int, n)
	noRobbery := make([]int, n)
	robbery[0] = nums[0]
	noRobbery[0] = 0
	for i := 1; i < n; i++ {
		robbery[i] = noRobbery[i-1] + nums[i]
		noRobbery[i] = max(noRobbery[i-1], robbery[i-1])
	}
	return max(robbery[n-1], noRobbery[n-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

