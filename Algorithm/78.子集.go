/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 *
 * https://leetcode-cn.com/problems/subsets/description/
 *
 * algorithms
 * Medium (72.99%)
 * Likes:    238
 * Dislikes: 0
 * Total Accepted:    19.6K
 * Total Submissions: 26.8K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
 *
 * 说明：解集不能包含重复的子集。
 *
 * 示例:
 *
 * 输入: nums = [1,2,3]
 * 输出:
 * [
 * ⁠ [3],
 * [1],
 * [2],
 * [1,2,3],
 * [1,3],
 * [2,3],
 * [1,2],
 * []
 * ]
 *
 */
func largestRectangleArea(heights []int) int {
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

