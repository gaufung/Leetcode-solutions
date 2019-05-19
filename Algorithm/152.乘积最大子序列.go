/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子序列
 *
 * https://leetcode-cn.com/problems/maximum-product-subarray/description/
 *
 * algorithms
 * Medium (33.77%)
 * Likes:    175
 * Dislikes: 0
 * Total Accepted:    9.8K
 * Total Submissions: 29.1K
 * Testcase Example:  '[2,3,-2,4]'
 *
 * 给定一个整数数组 nums ，找出一个序列中乘积最大的连续子序列（该序列至少包含一个数）。
 *
 * 示例 1:
 *
 * 输入: [2,3,-2,4]
 * 输出: 6
 * 解释: 子数组 [2,3] 有最大乘积 6。
 *
 *
 * 示例 2:
 *
 * 输入: [-2,0,-1]
 * 输出: 0
 * 解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
 *
 */
func maxProduct(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}
	dpMax := make([]int, n)
	dpMin := make([]int, n)
	dpMax[0], dpMin[0] = nums[0], nums[0]
	for i := 1; i < n; i++ {
		dpMax[i] = max(max(dpMax[i-1]*nums[i], dpMin[i-1]*nums[i]), nums[i])
		dpMin[i] = min(min(dpMax[i-1]*nums[i], dpMin[i-1]*nums[i]), nums[i])
	}
	max := dpMax[0]
	for i := 0; i < n; i++ {
		if dpMax[i] > max {
			max = dpMax[i]
		}
	}
	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

