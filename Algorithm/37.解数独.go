/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 *
 * https://leetcode-cn.com/problems/sudoku-solver/description/
 *
 * algorithms
 * Hard (53.79%)
 * Likes:    158
 * Dislikes: 0
 * Total Accepted:    6.3K
 * Total Submissions: 11.8K
 * Testcase Example:  '[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]'
 *
 * 编写一个程序，通过已填充的空格来解决数独问题。
 *
 * 一个数独的解法需遵循如下规则：
 *
 *
 * 数字 1-9 在每一行只能出现一次。
 * 数字 1-9 在每一列只能出现一次。
 * 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
 *
 *
 * 空白格用 '.' 表示。
 *
 *
 *
 * 一个数独。
 *
 *
 *
 * 答案被标成红色。
 *
 * Note:
 *
 *
 * 给定的数独序列只包含数字 1-9 和字符 '.' 。
 * 你可以假设给定的数独只有唯一解。
 * 给定数独永远是 9x9 形式的。
 *
 *
 */
func solveSudoku(board [][]byte) {
	rowUsed := make([][]bool, 9)
	for i := 0; i < 9; i++ {
		rowUsed[i] = make([]bool, 9)
	}
	colUsed := make([][]bool, 9)
	for i := 0; i < 9; i++ {
		colUsed[i] = make([]bool, 9)
	}
	blockUsed := make([][]bool, 9)
	for i := 0; i < 9; i++ {
		blockUsed[i] = make([]bool, 9)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				val := int(board[i][j] - '1')
				rowUsed[i][val] = true
				colUsed[j][val] = true
				blockUsed[i/3*3+j/3][val] = true
			}
		}
	}
	dfs(board, 0, rowUsed, colUsed, blockUsed)
}

func dfs(board [][]byte, pos int, rowUsed, colUsed, blockUsed [][]bool) bool {
	if pos >= 81 {
		return true
	}
	row := pos / 9
	col := pos % 9
	block := row/3*3 + col/3
	if board[row][col] != '.' {
		return dfs(board, pos+1, rowUsed, colUsed, blockUsed)
	}
	for i := 1; i < 10; i++ {
		if rowUsed[row][i-1] || colUsed[col][i-1] || blockUsed[block][i-1] {
			continue
		}
		board[row][col] = byte(i + '0')
		rowUsed[row][i-1] = true
		colUsed[col][i-1] = true
		blockUsed[block][i-1] = true
		if dfs(board, pos+1, rowUsed, colUsed, blockUsed) {
			return true
		} else {
			rowUsed[row][i-1] = false
			colUsed[col][i-1] = false
			blockUsed[block][i-1] = false
			board[row][col] = '.'
		}
	}
	return false
}

