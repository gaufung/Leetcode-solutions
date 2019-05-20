/*
 * @lc app=leetcode.cn id=149 lang=golang
 *
 * [149] 直线上最多的点数
 *
 * https://leetcode-cn.com/problems/max-points-on-a-line/description/
 *
 * algorithms
 * Hard (16.96%)
 * Likes:    56
 * Dislikes: 0
 * Total Accepted:    2.8K
 * Total Submissions: 16.3K
 * Testcase Example:  '[[1,1],[2,2],[3,3]]'
 *
 * 给定一个二维平面，平面上有 n 个点，求最多有多少个点在同一条直线上。
 *
 * 示例 1:
 *
 * 输入: [[1,1],[2,2],[3,3]]
 * 输出: 3
 * 解释:
 * ^
 * |
 * |        o
 * |     o
 * |  o
 * +------------->
 * 0  1  2  3  4
 *
 *
 * 示例 2:
 *
 * 输入: [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
 * 输出: 4
 * 解释:
 * ^
 * |
 * |  o
 * |     o        o
 * |        o
 * |  o        o
 * +------------------->
 * 0  1  2  3  4  5  6
 *
 */
func maxPoints(points [][]int) int {
	lines := make([]*Line, 0)
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			lines = append(lines, &Line{points[i][0], points[i][1], points[j][0], points[j][1]})
		}
	}
	longest := make([]int, len(lines))
	max := 0
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(lines); j++ {
			if lines[j].inLine(points[i][0], points[i][1]) {
				longest[j] += 1
				if longest[j] > max {
					max = longest[j]
				}
			}
		}
	}
	return max
}

type Line struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

func (l *Line) inLine(x, y int) bool {
	if l.X1 == l.X2 && l.X1 == x {
		return true
	}
	if l.Y1 == l.Y1 && l.Y1 == y {
		return true
	}
	return (y-l.Y1)*(l.X1-l.X2) == (x-l.X1)*(l.Y1-l.Y2)
}
