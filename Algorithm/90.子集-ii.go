import "sort"

/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	n := len(nums)
	if n <= 0 {
		return result
	}
	stack := NewStack()
	stack.Push(-1)
	for stack.Size() > 0 {
		top := stack.Pop() + 1
		if top >= n {
			continue
		}
		if top != 0 {
			if nums[top] == nums[top-1] {
				stack.Push(top)
				continue
			}
		}
		for i := top; i < n; i++ {
			stack.Push(i)
			result = append(result, buildStack(stack, nums))
		}
	}
	result = append(result, buildStack(stack, nums))
	return result
}

func buildStack(stack *Stack, nums []int) []int {
	result := make([]int, 0)
	for _, index := range stack.Values() {
		result = append(result, nums[index])
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

