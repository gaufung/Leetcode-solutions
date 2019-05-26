/*
 * @lc app=leetcode.cn id=491 lang=golang
 *
 * [491] 递增子序列
 *
 * https://leetcode-cn.com/problems/increasing-subsequences/description/
 *
 * algorithms
 * Medium (42.09%)
 * Likes:    37
 * Dislikes: 0
 * Total Accepted:    1.3K
 * Total Submissions: 3K
 * Testcase Example:  '[4,6,7,7]'
 *
 * 给定一个整型数组, 你的任务是找到所有该数组的递增子序列，递增子序列的长度至少是2。
 *
 * 示例:
 *
 *
 * 输入: [4, 6, 7, 7]
 * 输出: [[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7,7],
 * [4,7,7]]
 *
 * 说明:
 *
 *
 * 给定数组的长度不会超过15。
 * 数组中的整数范围是 [-100,100]。
 * 给定数组中可能包含重复数字，相等的数字应该被视为递增的一种情况。
 *
 *
 */
func findSubsequences(nums []int) [][]int {
	result := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		r := greaterSequence(nums[i], i+1, nums)
		for _, rr := range r {
			if len(rr) > 0 {
				result = append(result, append([]int{nums[i]}, rr...))
			}

		}
	}
	return result
}

func greaterSequence(base int, start int, nums []int) [][]int {
	if start >= len(nums) {
		return [][]int{
			[]int{},
		}
	}
	result := make([][]int, 0)
	for i := start; i < len(nums); i++ {
		if nums[i] >= base {
			val1 := []int{nums[i]}
			result1 := greaterSequence(nums[i], i+1, nums)
			for _, r1 := range result1 {
				result = append(result, append(val1, r1...))
			}
			val2 := []int{}
			result2 := greaterSequence(base, i+2, nums)
			for _, r2 := range result2 {
				result = append(result, append(val2, r2...))
			}
			break
		}
	}
	return result
}

