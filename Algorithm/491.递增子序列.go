/*
 * @lc app=leetcode.cn id=491 lang=golang
 *
 * [491] 递增子序列
 *
 * https://leetcode-cn.com/problems/increasing-subsequences/description/
 *
 * algorithms
 * Medium (42.69%)
 * Likes:    43
 * Dislikes: 0
 * Total Accepted:    1.5K
 * Total Submissions: 3.6K
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
	item := make([]int, 0)
	idx := make([]int, 0)
	res := make([][]int, 0)
	dfs(nums, 0, idx, item, &res)
	return res
}

func dfs(nums []int, pos int, idx []int, item []int, res *[][]int) {
	if len(item) > 1 {
		*res = append(*res, item)
	}
	for i := pos; i < len(nums); i++ {

		if len(idx) == 0 {
			idx1 := append(idx, i)
			item1 := append(item, nums[i])
			dfs(nums, i+1, idx1, item1, res)
		} else {
			lastIdx := idx[len(idx)-1]
			found := false
			for j := lastIdx + 1; j < i; j++ {
				if nums[i] == nums[j] {
					found = true
					break
				}
			}
			if found {
				continue
			}
			last := nums[lastIdx]
			if nums[i] >= last {
				idx2 := append(idx, i)
				item2 := append(item, nums[i])
				dfs(nums, i+1, idx2, item2, res)
			}
		}
	}
}

