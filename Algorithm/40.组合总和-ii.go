import "sort"

/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 *
 * https://leetcode-cn.com/problems/combination-sum-ii/description/
 *
 * algorithms
 * Medium (55.04%)
 * Likes:    102
 * Dislikes: 0
 * Total Accepted:    12.8K
 * Total Submissions: 23.2K
 * Testcase Example:  '[10,1,2,7,6,1,5]\n8'
 *
 * 给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 *
 * candidates 中的每个数字在每个组合中只能使用一次。
 *
 * 说明：
 *
 *
 * 所有数字（包括目标数）都是正整数。
 * 解集不能包含重复的组合。
 *
 *
 * 示例 1:
 *
 * 输入: candidates = [10,1,2,7,6,1,5], target = 8,
 * 所求解集为:
 * [
 * ⁠ [1, 7],
 * ⁠ [1, 2, 5],
 * ⁠ [2, 6],
 * ⁠ [1, 1, 6]
 * ]
 *
 *
 * 示例 2:
 *
 * 输入: candidates = [2,5,2,1,2], target = 5,
 * 所求解集为:
 * [
 * [1,2,2],
 * [5]
 * ]
 *
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

