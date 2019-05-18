/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
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

