/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */
func longestValidParentheses(s string) int {
	n := len(s)
	if n <= 0 {
		return 0
	}
	dp := make([]int, n)
	stack := NewStack()
	for i := 0; i < n; i++ {
		if stack.Size() == 0 || s[i] == '(' {
			stack.Push(i)
		} else {
			if s[stack.Top()] == '(' {
				dp[stack.Pop()] = 1
				dp[i] = 1
			} else {
				stack.Push(i)
			}
		}
	}
	val := dp[0]
	for i := 1; i < n; i++ {
		if dp[i] == 1 {
			dp[i] = dp[i-1] + 1
		}
		val = max(val, dp[i])
	}
	return val
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
	val := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return val
}

func (s *Stack) Top() int {
	return s.elements[len(s.elements)-1]
}

func (s *Stack) Size() int {
	return len(s.elements)
}
