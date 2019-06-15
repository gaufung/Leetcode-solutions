/*
 * @lc app=leetcode.cn id=417 lang=golang
 *
 * [417] 太平洋大西洋水流问题
 *
 * https://leetcode-cn.com/problems/pacific-atlantic-water-flow/description/
 *
 * algorithms
 * Medium (38.98%)
 * Likes:    34
 * Dislikes: 0
 * Total Accepted:    1.3K
 * Total Submissions: 3.3K
 * Testcase Example:  '[[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]'
 *
 * 给定一个 m x n 的非负整数矩阵来表示一片大陆上各个单元格的高度。“太平洋”处于大陆的左边界和上边界，而“大西洋”处于大陆的右边界和下边界。
 *
 * 规定水流只能按照上、下、左、右四个方向流动，且只能从高到低或者在同等高度上流动。
 *
 * 请找出那些水流既可以流动到“太平洋”，又能流动到“大西洋”的陆地单元的坐标。
 *
 *
 *
 * 提示：
 *
 *
 * 输出坐标的顺序不重要
 * m 和 n 都小于150
 *
 *
 *
 *
 * 示例：
 *
 *
 *
 *
 * 给定下面的 5x5 矩阵:
 *
 * ⁠ 太平洋 ~   ~   ~   ~   ~
 * ⁠      ~  1   2   2   3  (5) *
 * ⁠      ~  3   2   3  (4) (4) *
 * ⁠      ~  2   4  (5)  3   1  *
 * ⁠      ~ (6) (7)  1   4   5  *
 * ⁠      ~ (5)  1   1   2   4  *
 * ⁠         *   *   *   *   * 大西洋
 *
 * 返回:
 *
 * [[0, 4], [1, 3], [1, 4], [2, 2], [3, 0], [3, 1], [4, 0]] (上图中带括号的单元).
 *
 *
 *
 *
 */
func pacificAtlantic(matrix [][]int) [][]int {
	m := len(matrix)
	if m <= 0 {
		return [][]int{}
	}
	n := len(matrix[0])
	if n <= 0 {
		return [][]int{}
	}
	pacificVisited := make([][]bool, m)
	atlnaticVisited := make([][]bool, m)
	for i := 0; i < m; i++ {
		pacificVisited[i] = make([]bool, n)
		atlnaticVisited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		if pacificVisited[i][0] == false {
			pacificBack(matrix, i, 0, pacificVisited)
		}
		if atlnaticVisited[i][n-1] == false {
			atlanticBack(matrix, i, n-1, atlnaticVisited)
		}
	}
	for j := 0; j < n; j++ {
		if pacificVisited[0][j] == false {
			pacificBack(matrix, 0, j, pacificVisited)
		}
		if atlnaticVisited[m-1][j] == false {
			atlanticBack(matrix, m-1, j, atlnaticVisited)
		}
	}
	result := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pacificVisited[i][j] && atlnaticVisited[i][j] {
				result = append(result, []int{i, j})
			}
		}
	}
	return result

}

func pacificBack(matrix [][]int, i, j int, visited [][]bool) {
	m, n := len(matrix), len(matrix[0])
	visited[i][j] = true
	if i+1 < m && visited[i+1][j] == false && matrix[i+1][j] >= matrix[i][j] {
		pacificBack(matrix, i+1, j, visited)
	}
	if j+1 < n && visited[i][j+1] == false && matrix[i][j+1] >= matrix[i][j] {
		pacificBack(matrix, i, j+1, visited)
	}
	if i-1 >= 0 && visited[i-1][j] == false && matrix[i-1][j] >= matrix[i][j] {
		pacificBack(matrix, i-1, j, visited)
	}
	if j-1 >= 0 && visited[i][j-1] == false && matrix[i][j-1] >= matrix[i][j] {
		pacificBack(matrix, i, j-1, visited)
	}
}

func atlanticBack(matrix [][]int, i, j int, visited [][]bool) {
	m, n := len(matrix), len(matrix[0])
	visited[i][j] = true
	if i+1 < m && visited[i+1][j] == false && matrix[i+1][j] >= matrix[i][j] {
		atlanticBack(matrix, i+1, j, visited)
	}
	if j+1 < n && visited[i][j+1] == false && matrix[i][j+1] >= matrix[i][j] {
		atlanticBack(matrix, i, j+1, visited)
	}
	if i-1 >= 0 && visited[i-1][j] == false && matrix[i-1][j] >= matrix[i][j] {
		atlanticBack(matrix, i-1, j, visited)
	}
	if j-1 >= 0 && visited[i][j-1] == false && matrix[i][j-1] >= matrix[i][j] {
		atlanticBack(matrix, i, j-1, visited)
	}
}

// 1 2 3
// 8 9 4
// 7 6 5 
