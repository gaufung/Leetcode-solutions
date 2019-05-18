/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 *
 * https://leetcode-cn.com/problems/permutations/description/
 *
 * algorithms
 * Medium (68.85%)
 * Likes:    267
 * Dislikes: 0
 * Total Accepted:    24.5K
 * Total Submissions: 35.5K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个没有重复数字的序列，返回其所有可能的全排列。
 *
 * 示例:
 *
 * 输入: [1,2,3]
 * 输出:
 * [
 * ⁠ [1,2,3],
 * ⁠ [1,3,2],
 * ⁠ [2,1,3],
 * ⁠ [2,3,1],
 * ⁠ [3,1,2],
 * ⁠ [3,2,1]
 * ]
 *
 */
func permute(nums []int) [][]int {
	n := len(nums)
	if n <= 0 {
		return [][]int{}
	}
	if n == 1 {
		return [][]int{
			[]int{nums[0]},
		}
	}
	result := make([][]int, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		leftResult := permute(removeKth(nums, i))
		for _, items := range leftResult {
			result = append(result, append(items, nums[i]))
		}
	}
	return result
}

func removeKth(nums []int, k int) []int {
	result := make([]int, 0)
	for i, val := range nums {
		if i != k {
			result = append(result, val)
		}
	}
	return result
}

