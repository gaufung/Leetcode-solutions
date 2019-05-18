/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 *
 * https://leetcode-cn.com/problems/minimum-path-sum/description/
 *
 * algorithms
 * Medium (60.67%)
 * Likes:    201
 * Dislikes: 0
 * Total Accepted:    15.5K
 * Total Submissions: 25.6K
 * Testcase Example:  '[[1,3,1],[1,5,1],[4,2,1]]'
 *
 * 给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
 *
 * 说明：每次只能向下或者向右移动一步。
 *
 * 示例:
 *
 * 输入:
 * [
 * [1,3,1],
 * ⁠ [1,5,1],
 * ⁠ [4,2,1]
 * ]
 * 输出: 7
 * 解释: 因为路径 1→3→1→1→1 的总和最小。
 *
 *
 */
func minPathSum(grid [][]int) int {
	m := len(grid)
	if m <= 0 {
		return 0
	}
	n := len(grid[0])
	if n <= 0 {
		return 0
	}
	paths := make([][]int, m)
	for i := 0; i < m; i++ {
		paths[i] = make([]int, n)
	}
	paths[0][0] = grid[0][0]
	for j := 1; j < n; j++ {
		paths[0][j] = paths[0][j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		paths[i][0] = paths[i-1][0] + grid[i][0]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			paths[i][j] = min(paths[i-1][j], paths[i][j-1]) + grid[i][j]
		}
	}
	return paths[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

