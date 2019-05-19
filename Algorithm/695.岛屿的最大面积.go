/*
 * @lc app=leetcode.cn id=695 lang=golang
 *
 * [695] 岛屿的最大面积
 *
 * https://leetcode-cn.com/problems/max-area-of-island/description/
 *
 * algorithms
 * Medium (54.24%)
 * Likes:    116
 * Dislikes: 0
 * Total Accepted:    6.3K
 * Total Submissions: 11.7K
 * Testcase Example:  '[[1,1,0,0,0],[1,1,0,0,0],[0,0,0,1,1],[0,0,0,1,1]]'
 *
 * 给定一个包含了一些 0 和 1的非空二维数组 grid , 一个 岛屿 是由四个方向 (水平或垂直) 的 1 (代表土地)
 * 构成的组合。你可以假设二维矩阵的四个边缘都被水包围着。
 *
 * 找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为0。)
 *
 * 示例 1:
 *
 *
 * [[0,0,1,0,0,0,0,1,0,0,0,0,0],
 * ⁠[0,0,0,0,0,0,0,1,1,1,0,0,0],
 * ⁠[0,1,1,0,1,0,0,0,0,0,0,0,0],
 * ⁠[0,1,0,0,1,1,0,0,1,0,1,0,0],
 * ⁠[0,1,0,0,1,1,0,0,1,1,1,0,0],
 * ⁠[0,0,0,0,0,0,0,0,0,0,1,0,0],
 * ⁠[0,0,0,0,0,0,0,1,1,1,0,0,0],
 * ⁠[0,0,0,0,0,0,0,1,1,0,0,0,0]]
 *
 *
 * 对于上面这个给定矩阵应返回 6。注意答案不应该是11，因为岛屿只能包含水平或垂直的四个方向的‘1’。
 *
 * 示例 2:
 *
 *
 * [[0,0,0,0,0,0,0,0]]
 *
 * 对于上面这个给定的矩阵, 返回 0。
 *
 * 注意: 给定的矩阵grid 的长度和宽度都不超过 50。
 *
 */
func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	if m <= 0 {
		return 0
	}
	n := len(grid[0])
	if n <= 0 {
		return 0
	}
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	max := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && visited[i][j] == false {
				val := area(grid, visited, i, j, m, n)
				if val > max {
					max = val
				}
			}
		}
	}
	return max
}

func area(grid [][]int, visited [][]bool, i, j int, m, n int) int {
	visited[i][j] = true
	val := 1
	if i-1 >= 0 && visited[i-1][j] == false && grid[i-1][j] == 1 {
		val += area(grid, visited, i-1, j, m, n)
	}
	if i+1 < m && visited[i+1][j] == false && grid[i+1][j] == 1 {
		val += area(grid, visited, i+1, j, m, n)
	}
	if j-1 >= 0 && visited[i][j-1] == false && grid[i][j-1] == 1 {
		val += area(grid, visited, i, j-1, m, n)
	}
	if j+1 < n && visited[i][j+1] == false && grid[i][j+1] == 1 {
		val += area(grid, visited, i, j+1, m, n)
	}
	return val
}

