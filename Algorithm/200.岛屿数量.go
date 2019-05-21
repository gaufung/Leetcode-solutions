/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 *
 * https://leetcode-cn.com/problems/number-of-islands/description/
 *
 * algorithms
 * Medium (43.03%)
 * Likes:    153
 * Dislikes: 0
 * Total Accepted:    14.1K
 * Total Submissions: 32.7K
 * Testcase Example:  '[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]'
 *
 * 给定一个由 '1'（陆地）和
 * '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。
 *
 * 示例 1:
 *
 * 输入:
 * 11110
 * 11010
 * 11000
 * 00000
 *
 * 输出: 1
 *
 *
 * 示例 2:
 *
 * 输入:
 * 11000
 * 11000
 * 00100
 * 00011
 *
 * 输出: 3
 *
 *
 */
func numIslands(grid [][]byte) int {
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
	cnt := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && visited[i][j] == false {
				dfs(grid, i, j, m, n, visited)
				cnt++
			}
		}
	}
	return cnt

}

func dfs(grid [][]byte, i, j int, m, n int, visited [][]bool) {
	visited[i][j] = true
	if i-1 >= 0 && grid[i-1][j] == '1' && visited[i-1][j] == false {
		dfs(grid, i-1, j, m, n, visited)
	}
	if i+1 < m && grid[i+1][j] == '1' && visited[i+1][j] == false {
		dfs(grid, i+1, j, m, n, visited)
	}
	if j-1 >= 0 && grid[i][j-1] == '1' && visited[i][j-1] == false {
		dfs(grid, i, j-1, m, n, visited)
	}
	if j+1 < n && grid[i][j+1] == '1' && visited[i][j+1] == false {
		dfs(grid, i, j+1, m, n, visited)
	}
}

