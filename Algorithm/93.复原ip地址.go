import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原IP地址
 *
 * https://leetcode-cn.com/problems/restore-ip-addresses/description/
 *
 * algorithms
 * Medium (44.21%)
 * Likes:    94
 * Dislikes: 0
 * Total Accepted:    8.1K
 * Total Submissions: 18.4K
 * Testcase Example:  '"25525511135"'
 *
 * 给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
 *
 * 示例:
 *
 * 输入: "25525511135"
 * 输出: ["255.255.11.135", "255.255.111.35"]
 *
 */
func restoreIpAddresses(s string) []string {
	n := len(s)
	if n < 4 {
		return []string{}
	}
	stack := NewStack()
	stack.Push(0)
	stack.Push(0)
	result := make([]string, 0)
	for stack.Size() > 1 {
		top := stack.Pop() + 1
		if top > n {
			continue
		}
		for i := top; i <= n && stack.Size() < 6; i++ {
			if valid(s[stack.Top():i]) {
				stack.Push(i)
			}
		}
		if stack.Size() == 5 && stack.Top() == n {
			values := stack.Values()
			items := make([]string, 0)
			for i := 0; i < len(values)-1; i++ {
				items = append(items, s[values[i]:values[i+1]])
			}
			result = append(result, strings.Join(items, "."))
		}
	}
	// result := make([]string, 0)
	// for stack.Size() > 0 {
	// 	top := stack.Pop() + 1
	// 	if top >= n {
	// 		continue
	// 	}
	// 	if stack.Size() == 0 {
	// 		if valid(s[0:top]) {
	// 			stack.Push(top)
	// 		} else {
	// 			continue
	// 		}
	// 	} else {
	// 		if valid(s[stack.Top():top]) {
	// 			stack.Push(top)
	// 		} else {
	// 			continue
	// 		}
	// 	}
	// 	if stack.Size() == 1 {
	// 		index := stack.Top() + 1
	// 		for index < n {
	// 			if valid(s[stack.Top():index]) {
	// 				stack.Push(index)
	// 				break
	// 			}
	// 		}
	// 		if index == n {
	// 			continue
	// 		}
	// 	}
	// 	if stack.Size() == 2 {
	// 		index := stack.Top() + 1
	// 		for index < n {
	// 			if valid(s[stack.Top():index]) {
	// 				stack.Push(index)
	// 				break
	// 			}
	// 		}
	// 		if index == n {
	// 			continue
	// 		}
	// 	}
	// 	if stack.Size() == 3 {
	// 		if valid(s[stack.Top():n]) {
	// 			stack.Push(n)
	// 			items := make([]string, 0)
	// 			startIndex := 0
	// 			for _, index := range stack.Values() {
	// 				items = append(items, s[startIndex:index])
	// 				startIndex = index
	// 			}
	// 			result = append(result, strings.Join(items, "."))

	// 		} else {
	// 			continue
	// 		}
	// 	}
	// }

	return result
}

func valid(s string) bool {
	if s[0] == '0' && len(s) > 1 {
		return false
	}
	val, _ := strconv.Atoi(s)
	return val >= 0 && val <= 255
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

