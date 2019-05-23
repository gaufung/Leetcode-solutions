/*
 * @lc app=leetcode.cn id=223 lang=golang
 *
 * [223] 矩形面积
 *
 * https://leetcode-cn.com/problems/rectangle-area/description/
 *
 * algorithms
 * Medium (40.43%)
 * Likes:    22
 * Dislikes: 0
 * Total Accepted:    2.2K
 * Total Submissions: 5.4K
 * Testcase Example:  '-3\n0\n3\n4\n0\n-1\n9\n2'
 *
 * 在二维平面上计算出两个由直线构成的矩形重叠后形成的总面积。
 *
 * 每个矩形由其左下顶点和右上顶点坐标表示，如图所示。
 *
 *
 *
 * 示例:
 *
 * 输入: -3, 0, 3, 4, 0, -1, 9, 2
 * 输出: 45
 *
 * 说明: 假设矩形面积不会超出 int 的范围。
 *
 */
func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	if E >= C || G <= A {
		return (C-A)*(D-B) + (G-E)*(H-F)
	}
	if H <= B || F >= D {
		return (C-A)*(D-B) + (G-E)*(H-F)
	}
	a := max(A, E)
	b := max(B, F)
	c := min(C, G)
	d := min(D, H)
	return (C-A)*(D-B) + (G-E)*(H-F) - (c-a)*(d-b)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
