/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 *
 * https://leetcode-cn.com/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (35.30%)
 * Likes:    154
 * Dislikes: 0
 * Total Accepted:    13.9K
 * Total Submissions: 39.4K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
 *
 * 示例 1:
 *
 * 输入:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 4, 5, 6 ],
 * ⁠[ 7, 8, 9 ]
 * ]
 * 输出: [1,2,3,6,9,8,7,4,5]
 *
 *
 * 示例 2:
 *
 * 输入:
 * [
 * ⁠ [1, 2, 3, 4],
 * ⁠ [5, 6, 7, 8],
 * ⁠ [9,10,11,12]
 * ]
 * 输出: [1,2,3,4,8,12,11,10,9,5,6,7]
 *
 *
 */
func spiralOrder(matrix [][]int) []int {
	result := make([]int, 0)
	m := len(matrix)
	if m <= 0 {
		return result
	}
	n := len(matrix[0])
	if n <= 0 {
		return result
	}
	upper, bottom, left, right := 0, m, 0, n
	i, j := 0, 0
	isMore := true
	for isMore {
		isMore = false
		//top
		if i >= upper && i < bottom && j >= left && j < right {
			for k := j; k < right; k++ {
				isMore = true
				result = append(result, matrix[i][k])
			}
		}
		upper++
		i++
		j = right - 1
		//right
		if i >= upper && i < bottom && j >= left && j < right {
			for k := i; k < bottom; k++ {
				isMore = true
				result = append(result, matrix[k][j])
			}
		}
		right--
		j--
		i = bottom - 1
		//bottom
		if i >= upper && i < bottom && j >= left && j < right {
			for k := j; k >= left; k-- {
				isMore = true
				result = append(result, matrix[i][k])
			}
		}
		bottom--
		j = left
		i--
		//left
		if i >= upper && i < bottom && j >= left && j < right {
			for k := i; k >= upper; k-- {
				isMore = true
				result = append(result, matrix[k][j])
			}
		}
		left++
		j++
		i = upper
	}
	return result
}

