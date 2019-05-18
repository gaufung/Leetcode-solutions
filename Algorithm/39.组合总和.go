import "sort"

/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 *
 * https://leetcode-cn.com/problems/combination-sum/description/
 *
 * algorithms
 * Medium (64.42%)
 * Likes:    264
 * Dislikes: 0
 * Total Accepted:    18.3K
 * Total Submissions: 28.4K
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 *
 * candidates 中的数字可以无限制重复被选取。
 *
 * 说明：
 *
 *
 * 所有数字（包括 target）都是正整数。
 * 解集不能包含重复的组合。
 *
 *
 * 示例 1:
 *
 * 输入: candidates = [2,3,6,7], target = 7,
 * 所求解集为:
 * [
 * ⁠ [7],
 * ⁠ [2,2,3]
 * ]
 *
 *
 * 示例 2:
 *
 * 输入: candidates = [2,3,5], target = 8,
 * 所求解集为:
 * [
 * [2,2,2,2],
 * [2,3,3],
 * [3,5]
 * ]
 *
 */
func combinationSum(candidates []int, target int) [][]int {
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
		if top >= len(candidates) {
			continue
		}
		if leftValue-candidates[top] < 0 {
			continue
		}
		for leftValue > 0 {
			leftValue -= candidates[top]
			stack.Push(top)
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

