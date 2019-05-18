/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子序列
 */
func maxProduct(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}
	maxes := make([]int, n)
	mins := make([]int, n)
	maxes[0] = nums[0]
	mins[0] = nums[0]
	result := nums[0]
	for i := 1; i < n; i++ {
		maxes[i] = max(maxes[i-1]*nums[i], max(mins[i-1]*nums[i], nums[i]))
		mins[i] = min(maxes[i-1]*nums[i], min(mins[i-1]*nums[i], nums[i]))
		result = max(result, maxes[i])
	}
	return result
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
