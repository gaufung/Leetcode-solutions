/*
 * @lc app=leetcode.cn id=85 lang=golang
 *
 * [85] 最大矩形
 *
 * https://leetcode-cn.com/problems/maximal-rectangle/description/
 *
 * algorithms
 * Hard (42.93%)
 * Likes:    133
 * Dislikes: 0
 * Total Accepted:    3.9K
 * Total Submissions: 9K
 * Testcase Example:  '[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]'
 *
 * 给定一个仅包含 0 和 1 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
 *
 * 示例:
 *
 * 输入:
 * [
 * ⁠ ["1","0","1","0","0"],
 * ⁠ ["1","0","1","1","1"],
 * ⁠ ["1","1","1","1","1"],
 * ⁠ ["1","0","0","1","0"]
 * ]
 * 输出: 6
 *
 */
func maximalRectangle(matrix [][]byte) int {
	m := len(matrix)
	if m <= 0 {
		return 0
	}
	n := len(matrix[0])
	if n <= 0 {
		return 0
	}
	result := 0
	for i := 0; i < m; i++ {
		heights := createHeight(matrix, i)
		result = max(result, maxArea(heights))
	}
	return result

}

func createHeight(matrix [][]byte, row int) []int {
	m := len(matrix)
	n := len(matrix[0])
	heights := make([]int, n)
	for i := 0; i < n; i++ {
		val := 0
		for j := row; j < m; j++ {
			if matrix[j][i] == '0' {
				break
			} else {
				val += 1
			}
		}
		heights[i] = val
	}
	return heights
}

func maxArea(heights []int) int {
	n := len(heights)
	if n <= 0 {
		return 0
	}
	heights = append(heights, -1)
	result := 0
	stack := NewStack()
	for index, height := range heights {
		if stack.Size() == 0 || height >= heights[stack.Top()] {
			stack.Push(index)
		} else {
			for stack.Size() > 0 && height <= heights[stack.Top()] {
				h := heights[stack.Pop()]
				left := -1
				if stack.Size() > 0 {
					left = stack.Top()
				}
				result = max(result, h*(index-left-1))
			}
			stack.Push(index)
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Stack struct {
	elements []int
}

func NewStack() *Stack {
	return &Stack{
		elements: make([]int, 0),
	}
}

func (s *Stack) Push(val int) {
	s.elements = append(s.elements, val)
}

func (s *Stack) Pop() int {
	val := s.Top()
	s.elements = s.elements[:len(s.elements)-1]
	return val
}

func (s *Stack) Size() int {
	return len(s.elements)
}

func (s *Stack) Top() int {
	return s.elements[len(s.elements)-1]
}
func (s *Stack) Values() []int {
	return s.elements
}

