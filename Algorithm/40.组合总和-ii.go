import "sort"

/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	stack := NewStack()
	stack.Push(-1)
	leftValue := target
	result := make([][]int, 0)
	for stack.Size() > 0 {
		top := stack.Pop()
		if top != -1 {
			leftValue += candidates[top]
		}
		top++
		for top > 0 && top < len(candidates) && candidates[top] == candidates[top-1] {
			top++
		}
		if top >= len(candidates) {
			continue
		}
		if leftValue-candidates[top] < 0 {
			continue
		}
		for i := top; i < len(candidates) && leftValue > 0; i++ {
			leftValue -= candidates[i]
			stack.Push(i)
		}
		if leftValue == 0 {
			indexes := stack.Values()
			items := make([]int, len(indexes))
			for i, index := range indexes {
				items[i] = candidates[index]
			}
			result = append(result, items)
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