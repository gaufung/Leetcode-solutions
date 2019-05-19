/*
 * @lc app=leetcode.cn id=216 lang=golang
 *
 * [216] 组合总和 III
 *
 * https://leetcode-cn.com/problems/combination-sum-iii/description/
 *
 * algorithms
 * Medium (66.40%)
 * Likes:    48
 * Dislikes: 0
 * Total Accepted:    4.4K
 * Total Submissions: 6.6K
 * Testcase Example:  '3\n7'
 *
 * 找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。
 *
 * 说明：
 *
 *
 * 所有数字都是正整数。
 * 解集不能包含重复的组合。
 *
 *
 * 示例 1:
 *
 * 输入: k = 3, n = 7
 * 输出: [[1,2,4]]
 *
 *
 * 示例 2:
 *
 * 输入: k = 3, n = 9
 * 输出: [[1,2,6], [1,3,5], [2,3,4]]
 *
 *
 */
func combinationSum3(k int, n int) [][]int {
	result := make([][]int, 0)
	if n > k*9 || n < k*1 {
		return result
	}
	stack := NewStack()
	stack.Push(0)
	for stack.Size() > 0 {
		if stack.Size() > k {
			n += stack.Pop()
			continue
		} else {
			top := stack.Pop()
			n += top
			top++
			if top > 9 {
				continue
			}
			for i := top; i < 10 && n > 0 && stack.Size() < k; i++ {
				stack.Push(i)
				n -= i
			}
			if n == 0 && stack.Size() == k {
				value := make([]int, stack.Size())
				for i, val := range stack.Values() {
					value[i] = val
				}
				result = append(result, value)
			}
		}

	}
	return result
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

