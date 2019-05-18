/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 *
 * https://leetcode-cn.com/problems/surrounded-regions/description/
 *
 * algorithms
 * Medium (35.89%)
 * Likes:    64
 * Dislikes: 0
 * Total Accepted:    5K
 * Total Submissions: 14.1K
 * Testcase Example:  '[["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]'
 *
 * 给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。
 *
 * 找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
 *
 * 示例:
 *
 * X X X X
 * X O O X
 * X X O X
 * X O X X
 *
 *
 * 运行你的函数后，矩阵变为：
 *
 * X X X X
 * X X X X
 * X X X X
 * X O X X
 *
 *
 * 解释:
 *
 * 被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O'
 * 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
 *
 */
func solve(board [][]byte) {
	m := len(board)
	if m <= 0 {
		return
	}
	n := len(board[0])
	if n <= 0 {
		return
	}
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		DFS(board, visited, i, 0, m, n)
		DFS(board, visited, i, n-1, m, n)
	}
	for j := 0; j < n; j++ {
		DFS(board, visited, 0, j, m, n)
		DFS(board, visited, m-1, j, m, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if visited[i][j] == false {
				board[i][j] = 'X'
			}
		}
	}
}

func DFS(board [][]byte, visited [][]bool, i, j int, m, n int) {
	if board[i][j] == 'X' || visited[i][j] == true {
		return
	}
	visited[i][j] = true
	if i-1 >= 0 && board[i-1][j] == 'O' && visited[i-1][j] == false {
		DFS(board, visited, i-1, j, m, n)
	}
	if i+1 < m && board[i+1][j] == 'O' && visited[i+1][j] == false {
		DFS(board, visited, i+1, j, m, n)
	}
	if j-1 >= 0 && board[i][j-1] == 'O' && visited[i][j-1] == false {
		DFS(board, visited, i, j-1, m, n)
	}
	if j+1 < n && board[i][j+1] == 'O' && visited[i][j+1] == false {
		DFS(board, visited, i, j+1, m, n)
	}
}

