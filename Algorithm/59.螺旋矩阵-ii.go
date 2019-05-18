/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 *
 * https://leetcode-cn.com/problems/spiral-matrix-ii/description/
 *
 * algorithms
 * Medium (72.53%)
 * Likes:    98
 * Dislikes: 0
 * Total Accepted:    9.5K
 * Total Submissions: 13.2K
 * Testcase Example:  '3'
 *
 * 给定一个正整数 n，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。
 *
 * 示例:
 *
 * 输入: 3
 * 输出:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 8, 9, 4 ],
 * ⁠[ 7, 6, 5 ]
 * ]
 *
 */
func generateMatrix(n int) [][]int {
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, n)
	}
	cnt := 0
	top, left, bottom, right := 0, 0, n-1, n-1
	for {
		for i := left; i <= right; i++ {
			cnt++
			mat[top][i] = cnt
		}
		top++
		for i := top; i <= bottom; i++ {
			cnt++
			mat[i][right] = cnt
		}
		right--
		for i := right; i >= left; i-- {
			cnt++
			mat[bottom][i] = cnt
		}
		bottom--
		for i := bottom; i >= top; i-- {
			cnt++
			mat[i][left] = cnt
		}
		left++
		if left > right && top > bottom {
			break
		}
	}
	return mat
}
