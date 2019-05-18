/*
 * @lc app=leetcode.cn id=52 lang=golang
 *
 * [52] N皇后 II
 */
func totalNQueens(n int) int {
	result := 0
	stack := NewStack()
	stack.Push(-1)
	for stack.Size() > 0 {
		top := stack.Pop() + 1
		if top >= n {
			continue
		}
		if valid(stack, top) {
			stack.Push(top)
		} else {
			stack.Push(top)
			continue
		}
		for i := stack.Size(); i < n; i++ {
			j := 0
			for ; j < n; j++ {
				if valid(stack, j) {
					stack.Push(j)
					break
				}
			}
			if j == n {
				break
			}
		}
		if stack.Size() == n {
			result++
		}
	}
	return result
}

func valid(s *Stack, pos int) bool {
	size := s.Size()
	for i, val := range s.Values() {
		if pos == val {
			return false
		}
		if abs(size-i) == abs(val-pos) {
			return false
		}
	}
	return true
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
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