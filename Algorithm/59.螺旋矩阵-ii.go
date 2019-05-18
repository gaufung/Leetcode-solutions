/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
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

