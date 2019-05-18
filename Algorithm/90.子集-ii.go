import "sort"

/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 *
 * https://leetcode-cn.com/problems/subsets-ii/description/
 *
 * algorithms
 * Medium (54.38%)
 * Likes:    85
 * Dislikes: 0
 * Total Accepted:    7K
 * Total Submissions: 12.8K
 * Testcase Example:  '[1,2,2]'
 *
 * 给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
 *
 * 说明：解集不能包含重复的子集。
 *
 * 示例:
 *
 * 输入: [1,2,2]
 * 输出:
 * [
 * ⁠ [2],
 * ⁠ [1],
 * ⁠ [1,2,2],
 * ⁠ [2,2],
 * ⁠ [1,2],
 * ⁠ []
 * ]
 *
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

