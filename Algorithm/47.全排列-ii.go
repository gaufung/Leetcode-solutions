import "sort"

/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 *
 * https://leetcode-cn.com/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (52.11%)
 * Likes:    100
 * Dislikes: 0
 * Total Accepted:    12K
 * Total Submissions: 23K
 * Testcase Example:  '[1,1,2]'
 *
 * 给定一个可包含重复数字的序列，返回所有不重复的全排列。
 *
 * 示例:
 *
 * 输入: [1,1,2]
 * 输出:
 * [
 * ⁠ [1,1,2],
 * ⁠ [1,2,1],
 * ⁠ [2,1,1]
 * ]
 *
 */
func permuteUnique(nums []int) [][]int {
	n := len(nums)
	if n <= 0 {
		return [][]int{}
	}
	if n == 1 {
		return [][]int{
			[]int{nums[0]},
		}
	}
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		if i != len(nums)-1 && nums[i] == nums[i+1] {
			continue
		}
		leftResult := permuteUnique(removeKth(nums, i))
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


