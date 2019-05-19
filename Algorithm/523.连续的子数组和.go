/*
 * @lc app=leetcode.cn id=523 lang=golang
 *
 * [523] 连续的子数组和
 *
 * https://leetcode-cn.com/problems/continuous-subarray-sum/description/
 *
 * algorithms
 * Medium (22.88%)
 * Likes:    32
 * Dislikes: 0
 * Total Accepted:    1.9K
 * Total Submissions: 8.5K
 * Testcase Example:  '[23,2,4,6,7]\n6'
 *
 * 给定一个包含非负数的数组和一个目标整数 k，编写一个函数来判断该数组是否含有连续的子数组，其大小至少为 2，总和为 k 的倍数，即总和为 n*k，其中
 * n 也是一个整数。
 *
 * 示例 1:
 *
 * 输入: [23,2,4,6,7], k = 6
 * 输出: True
 * 解释: [2,4] 是一个大小为 2 的子数组，并且和为 6。
 *
 *
 * 示例 2:
 *
 * 输入: [23,2,6,4,7], k = 6
 * 输出: True
 * 解释: [23,2,6,4,7]是大小为 5 的子数组，并且和为 42。
 *
 *
 * 说明:
 *
 *
 * 数组的长度不会超过10,000。
 * 你可以认为所有数字总和在 32 位有符号整数范围内。
 *
 *
 */
func checkSubarraySum(nums []int, k int) bool {
	n := len(nums)
	partialSum := make([]int, n)
	partialSum[0] = nums[0]
	for i := 1; i < n; i++ {
		partialSum[i] = partialSum[i-1] + nums[i]
	}
	for i := 1; i < n; i++ {
		if k == 0 {
			if partialSum[i] == 0 {
				return true
			}
			for j := i - 2; j >= 0; j-- {
				if partialSum[i]-partialSum[j] == 0 {
					return true
				}
			}
		} else {
			if partialSum[i]%k == 0 {
				return true
			}
			for j := i - 2; j >= 0; j-- {
				if (partialSum[i]-partialSum[j])%k == 0 {
					return true
				}
			}
		}

	}
	return false

}


