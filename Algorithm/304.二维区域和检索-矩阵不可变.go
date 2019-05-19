/*
 * @lc app=leetcode.cn id=304 lang=golang
 *
 * [304] 二维区域和检索 - 矩阵不可变
 *
 * https://leetcode-cn.com/problems/range-sum-query-2d-immutable/description/
 *
 * algorithms
 * Medium (39.38%)
 * Likes:    29
 * Dislikes: 0
 * Total Accepted:    1.7K
 * Total Submissions: 4.4K
 * Testcase Example:  '["NumMatrix","sumRegion","sumRegion","sumRegion"]\n[[[[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]],[2,1,4,3],[1,1,2,2],[1,2,2,4]]'
 *
 * 给定一个二维矩阵，计算其子矩形范围内元素的总和，该子矩阵的左上角为 (row1, col1) ，右下角为 (row2, col2)。
 *
 *
 * 上图子矩阵左上角 (row1, col1) = (2, 1) ，右下角(row2, col2) = (4, 3)，该子矩形内元素的总和为 8。
 *
 * 示例:
 *
 * 给定 matrix = [
 * ⁠ [3, 0, 1, 4, 2],
 * ⁠ [5, 6, 3, 2, 1],
 * ⁠ [1, 2, 0, 1, 5],
 * ⁠ [4, 1, 0, 1, 7],
 * ⁠ [1, 0, 3, 0, 5]
 * ]
 *
 * sumRegion(2, 1, 4, 3) -> 8
 * sumRegion(1, 1, 2, 2) -> 11
 * sumRegion(1, 2, 2, 4) -> 12
 *
 *
 * 说明:
 *
 *
 * 你可以假设矩阵不可变。
 * 会多次调用 sumRegion 方法。
 * 你可以假设 row1 ≤ row2 且 col1 ≤ col2。
 *
 *
 */
type NumMatrix struct {
	M         int
	N         int
	SubMatrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	if m <= 0 {
		return NumMatrix{}
	}
	n := len(matrix[0])
	if n <= 0 {
		return NumMatrix{}
	}
	subMatrix := make([][]int, m)
	for i := 0; i < m; i++ {
		subMatrix[i] = make([]int, n)
	}
	subMatrix[0][0] = matrix[0][0]
	for j := 1; j < n; j++ {
		subMatrix[0][j] = subMatrix[0][j-1] + matrix[0][j]
	}
	for i := 1; i < m; i++ {
		subMatrix[i][0] = subMatrix[i-1][0] + matrix[i][0]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			subMatrix[i][j] = subMatrix[i-1][j] + subMatrix[i][j-1] - subMatrix[i-1][j-1] + matrix[i][j]
		}
	}
	return NumMatrix{
		M:         m,
		N:         n,
		SubMatrix: subMatrix,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if row1 == 0 {
		if col1 == 0 {
			return this.SubMatrix[row2][col2]
		} else {
			return this.SubMatrix[row2][col2] - this.SubMatrix[row2][col1-1]
		}
	} else {
		if col1 == 0 {
			return this.SubMatrix[row2][col2] - this.SubMatrix[row1-1][col2]
		} else {
			return this.SubMatrix[row2][col2] - this.SubMatrix[row1-1][col2] - this.SubMatrix[row2][col1-1] + this.SubMatrix[row1-1][col1-1]
		}
	}
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

