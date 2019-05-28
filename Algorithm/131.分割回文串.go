/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 *
 * https://leetcode-cn.com/problems/palindrome-partitioning/description/
 *
 * algorithms
 * Medium (62.04%)
 * Likes:    110
 * Dislikes: 0
 * Total Accepted:    7.2K
 * Total Submissions: 11.5K
 * Testcase Example:  '"aab"'
 *
 * 给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
 *
 * 返回 s 所有可能的分割方案。
 *
 * 示例:
 *
 * 输入: "aab"
 * 输出:
 * [
 * ⁠ ["aa","b"],
 * ⁠ ["a","a","b"]
 * ]
 *
 */
func partition(s string) [][]string {
	n := len(s)
	stack := NewStack()
	stack.Push(0)
	result := make([][]string, 0)
	for stack.Size() > 0 {
		top := stack.Pop() + 1
		if top > n {
			continue
		} else {
			if stack.Size() == 0 {
				stack.Push(top)
				if !Valid(s[0:top]) {
					continue
				}
			} else {
				if !Valid(s[stack.Top():top]) {
					stack.Push(top)
					continue
				} else {
					stack.Push(top)
				}
			}
		}
		for i := stack.Top() + 1; i <= n; i++ {
			if Valid(s[stack.Top():i]) {
				stack.Push(i)
			}
		}
		if stack.Top() == n {
			// values := stack.Values()
			// indexes := append([]int{0}, values...)
			// items := make([]string, 0)
			// for i := 0; i < len(indexes)-1; i++ {
			// 	items = append(items, s[i:i+1])
			// }
			// result = append(result, items)
			values := stack.Values()
			items := []string{s[0:values[0]]}
			for i := 0; i < len(values)-1; i++ {
				items = append(items, s[values[i]:values[i+1]])
			}
			result = append(result, items)
		}
	}
	return result
}

func Valid(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
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
