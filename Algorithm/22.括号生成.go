/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 *
 * https://leetcode-cn.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (70.10%)
 * Likes:    373
 * Dislikes: 0
 * Total Accepted:    22.8K
 * Total Submissions: 32.6K
 * Testcase Example:  '3'
 *
 * 给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。
 *
 * 例如，给出 n = 3，生成结果为：
 *
 * [
 * ⁠ "((()))",
 * ⁠ "(()())",
 * ⁠ "(())()",
 * ⁠ "()(())",
 * ⁠ "()()()"
 * ]
 *
 *
 */
func generateParenthesis(n int) []string {
	result := make([]string, 0)
	if n <= 0 {
		return result
	}
	stack := NewStack()
	stack.Push('?')
	leftCnt, rightCnt := 0, 0
	for stack.Size() > 0 {
		if stack.Top() == '?' { // init
			stack.Pop()
			stack.Push('(')
			leftCnt++
		} else if stack.Top() == '(' { //backtrace
			stack.Pop()
			leftCnt--
			stack.Push(')')
			rightCnt++
			if rightCnt > leftCnt {
				stack.Pop()
				rightCnt--
				continue
			}
		} else {
			stack.Pop()
			rightCnt--
			continue
		}
		for i := leftCnt; i < n; i++ {
			stack.Push('(')
			leftCnt++
		}
		for i := rightCnt; i < n; i++ {
			stack.Push(')')
			rightCnt++
		}
		result = append(result, string(stack.Values()))
	}
	return result
}

type Stack struct {
	elements []byte
}

func NewStack() *Stack {
	return &Stack{
		elements: make([]byte, 0),
	}
}

func (s *Stack) Push(val byte) {
	s.elements = append(s.elements, val)
}

func (s *Stack) Top() byte {
	return s.elements[len(s.elements)-1]
}
func (s *Stack) Pop() {
	s.elements = s.elements[:len(s.elements)-1]
}
func (s *Stack) Size() int {
	return len(s.elements)
}
func (s *Stack) Values() []byte {
	return s.elements
}



