/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 *
 * https://leetcode-cn.com/problems/combinations/description/
 *
 * algorithms
 * Medium (67.44%)
 * Likes:    120
 * Dislikes: 0
 * Total Accepted:    10.9K
 * Total Submissions: 16.2K
 * Testcase Example:  '4\n2'
 *
 * 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
 *
 * 示例:
 *
 * 输入: n = 4, k = 2
 * 输出:
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 *
 */
func combine(n int, k int) [][]int {
	return comb(1, n, k)
}

func comb(from, to int, k int) [][]int {
	if k == 0 {
		return [][]int{}
	}
	result := make([][]int, 0)
	if k == 1 {
		for i := from; i <= to; i++ {
			result = append(result, []int{i})
		}
		return result
	}
	for i := from; i <= to-k+1; i++ {
		leftResult := comb(i+1, to, k-1)
		for _, item := range leftResult {
			result = append(result, append(item, i))
		}
	}
	return result
}

