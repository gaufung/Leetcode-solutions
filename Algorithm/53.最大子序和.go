/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 *
 * https://leetcode-cn.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (44.94%)
 * Likes:    924
 * Dislikes: 0
 * Total Accepted:    56.8K
 * Total Submissions: 126.3K
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
 *
 * 示例:
 *
 * 输入: [-2,1,-3,4,-1,2,1,-5,4],
 * 输出: 6
 * 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
 *
 *
 * 进阶:
 *
 * 如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
 *
 */
func maxSubArray(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	result := nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		result = max(result, dp[i])
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

