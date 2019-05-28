/*
 * @lc app=leetcode.cn id=498 lang=golang
 *
 * [498] 对角线遍历
 *
 * https://leetcode-cn.com/problems/diagonal-traverse/description/
 *
 * algorithms
 * Medium (37.79%)
 * Likes:    39
 * Dislikes: 0
 * Total Accepted:    3.6K
 * Total Submissions: 9.6K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给定一个含有 M x N 个元素的矩阵（M 行，N 列），请以对角线遍历的顺序返回这个矩阵中的所有元素，对角线遍历如下图所示。
 *
 *
 *
 * 示例:
 *
 * 输入:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 4, 5, 6 ],
 * ⁠[ 7, 8, 9 ]
 * ]
 *
 * 输出:  [1,2,4,7,5,3,6,8,9]
 *
 * 解释:
 *
 *
 *
 *
 *
 * 说明:
 *
 *
 * 给定矩阵中的元素总数不会超过 100000 。
 *
 *
 */
func findDiagonalOrder(matrix [][]int) []int {
	result := make([]int, 0)
	m := len(matrix)
	if m <= 0 {
		return result
	}
	n := len(matrix[0])
	if n <= 0 {
		return result
	}

	isUp := true
	for interp := 0; interp < m+n; interp++ {
		// i = - j + interp
		if isUp {
			for i := interp; i >= 0; i-- {
				j := interp - i
				if i < m && j < n {
					result = append(result, matrix[i][j])
				}
			}
			isUp = !isUp
		} else {
			for j := interp; j >= 0; j-- {
				i := interp - j
				if i < m && j < n {
					result = append(result, matrix[i][j])
				}
			}
			isUp = !isUp
		}
	}
	return result
}

