/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 *
 * https://leetcode-cn.com/problems/subsets/description/
 *
 * algorithms
 * Medium (73.17%)
 * Likes:    250
 * Dislikes: 0
 * Total Accepted:    20.4K
 * Total Submissions: 27.8K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
 *
 * 说明：解集不能包含重复的子集。
 *
 * 示例:
 *
 * 输入: nums = [1,2,3]
 * 输出:
 * [
 * ⁠ [3],
 * [1],
 * [2],
 * [1,2,3],
 * [1,3],
 * [2,3],
 * [1,2],
 * []
 * ]
 *
 */
func subsets(nums []int) [][]int {
	n := len(nums)
	if n == 0 {
		return [][]int{
			[]int{},
		}
	}
	if n == 1 {
		return [][]int{
			[]int{},
			[]int{nums[0]},
		}
	}
	result := make([][]int, 0)
	leftResult := subsets(nums[1:])
	for _, left := range leftResult {
		result = append(result, left)
	}
	for _, left := range leftResult {
		result = append(result, append([]int{nums[0]}, left...))
	}
	return result
}

