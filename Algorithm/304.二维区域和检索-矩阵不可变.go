/*
 * @lc app=leetcode.cn id=304 lang=golang
 *
 * [304] 二维区域和检索 - 矩阵不可变
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

