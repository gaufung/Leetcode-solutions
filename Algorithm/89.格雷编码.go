/*
 * @lc app=leetcode.cn id=89 lang=golang
 *
 * [89] 格雷编码
 */
func grayCode(n int) []int {
	result := make([]int, 1<<uint(n))
	for i := 0; i < (1 << uint(n)); i++ {
		result[i] = i ^ (i / 2)
	}
	return result
}

