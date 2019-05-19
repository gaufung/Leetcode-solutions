/*
 * @lc app=leetcode.cn id=329 lang=golang
 *
 * [329] 矩阵中的最长递增路径
 *
 * https://leetcode-cn.com/problems/longest-increasing-path-in-a-matrix/description/
 *
 * algorithms
 * Hard (37.33%)
 * Likes:    50
 * Dislikes: 0
 * Total Accepted:    2.3K
 * Total Submissions: 6.2K
 * Testcase Example:  '[[9,9,4],[6,6,8],[2,1,1]]'
 *
 * 给定一个整数矩阵，找出最长递增路径的长度。
 *
 * 对于每个单元格，你可以往上，下，左，右四个方向移动。 你不能在对角线方向上移动或移动到边界外（即不允许环绕）。
 *
 * 示例 1:
 *
 * 输入: nums =
 * [
 * ⁠ [9,9,4],
 * ⁠ [6,6,8],
 * ⁠ [2,1,1]
 * ]
 * 输出: 4
 * 解释: 最长递增路径为 [1, 2, 6, 9]。
 *
 * 示例 2:
 *
 * 输入: nums =
 * [
 * ⁠ [3,4,5],
 * ⁠ [3,2,6],
 * ⁠ [2,2,1]
 * ]
 * 输出: 4
 * 解释: 最长递增路径是 [3, 4, 5, 6]。注意不允许在对角线方向上移动。
 *
 *
 */
func longestIncreasingPath(matrix [][]int) int {
	m := len(matrix)
	if m <= 0 {
		return 0
	}
	n := len(matrix[0])
	if n <= 0 {
		return 0
	}
	mm := make(map[Pos]int, 0)
	max := 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			pos := Pos{i, j}
			val := increase(matrix, pos, mm)
			if val > max {
				max = val
			}
		}
	}
	return max
}

func increase(matrix [][]int, seed Pos, mm map[Pos]int) int {
	m := len(matrix)
	n := len(matrix[0])
	max := 0
	if seed.X-1 >= 0 && matrix[seed.X-1][seed.Y] > matrix[seed.X][seed.Y] {
		nextPos := Pos{seed.X - 1, seed.Y}
		if val, ok := mm[nextPos]; ok {
			if val > max {
				max = val
			}
		} else {
			val = increase(matrix, nextPos, mm)
			if val > max {
				max = val
			}
		}
	}
	if seed.X+1 < m && matrix[seed.X+1][seed.Y] > matrix[seed.X][seed.Y] {
		nextPos := Pos{seed.X + 1, seed.Y}
		if val, ok := mm[nextPos]; ok {
			if val > max {
				max = val
			}
		} else {
			val = increase(matrix, nextPos, mm)
			if val > max {
				max = val
			}
		}
	}
	if seed.Y-1 >= 0 && matrix[seed.X][seed.Y-1] > matrix[seed.X][seed.Y] {
		nextPos := Pos{seed.X, seed.Y - 1}
		if val, ok := mm[nextPos]; ok {
			if val > max {
				max = val
			}
		} else {
			val = increase(matrix, nextPos, mm)
			if val > max {
				max = val
			}
		}
	}
	if seed.Y+1 < n && matrix[seed.X][seed.Y+1] > matrix[seed.X][seed.Y] {
		nextPos := Pos{seed.X, seed.Y + 1}
		if val, ok := mm[nextPos]; ok {
			if val > max {
				max = val
			}
		} else {
			val = increase(matrix, nextPos, mm)
			if val > max {
				max = val
			}
		}
	}
	mm[seed] = 1 + max
	return mm[seed]
}

type Pos struct {
	X int
	Y int
}

