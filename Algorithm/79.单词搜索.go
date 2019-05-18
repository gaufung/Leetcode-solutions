/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 *
 * https://leetcode-cn.com/problems/word-search/description/
 *
 * algorithms
 * Medium (37.31%)
 * Likes:    134
 * Dislikes: 0
 * Total Accepted:    10.1K
 * Total Submissions: 27K
 * Testcase Example:  '[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"'
 *
 * 给定一个二维网格和一个单词，找出该单词是否存在于网格中。
 *
 * 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
 *
 * 示例:
 *
 * board =
 * [
 * ⁠ ['A','B','C','E'],
 * ⁠ ['S','F','C','S'],
 * ⁠ ['A','D','E','E']
 * ]
 *
 * 给定 word = "ABCCED", 返回 true.
 * 给定 word = "SEE", 返回 true.
 * 给定 word = "ABCB", 返回 false.
 *
 */
func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(board, visited, word, i, j, 0, m, n) == true {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, visited [][]bool, word string, i, j int, pos, m, n int) bool {
	if pos == len(word) {
		return true
	}
	if i < 0 || i >= m || j < 0 || j >= n || word[pos] != board[i][j] || visited[i][j] == true {
		return false
	}
	visited[i][j] = true
	found := dfs(board, visited, word, i-1, j, pos+1, m, n) || dfs(board, visited, word, i+1, j, pos+1, m, n) || dfs(board, visited, word, i, j-1, pos+1, m, n) || dfs(board, visited, word, i, j+1, pos+1, m, n)
	visited[i][j] = false
	return found
}

