/*
 * @lc app=leetcode.cn id=385 lang=golang
 *
 * [385] 迷你语法分析器
 *
 * https://leetcode-cn.com/problems/mini-parser/description/
 *
 * algorithms
 * Medium (38.86%)
 * Likes:    9
 * Dislikes: 0
 * Total Accepted:    560
 * Total Submissions: 1.4K
 * Testcase Example:  '"324"'
 *
 * 给定一个用字符串表示的整数的嵌套列表，实现一个解析它的语法分析器。
 *
 * 列表中的每个元素只可能是整数或整数嵌套列表
 *
 * 提示：你可以假定这些字符串都是格式良好的：
 *
 *
 * 字符串非空
 * 字符串不包含空格
 * 字符串只包含数字0-9, [, - ,, ]
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 给定 s = "324",
 *
 * 你应该返回一个 NestedInteger 对象，其中只包含整数值 324。
 *
 *
 *
 *
 * 示例 2：
 *
 *
 * 给定 s = "[123,[456,[789]]]",
 *
 * 返回一个 NestedInteger 对象包含一个有两个元素的嵌套列表：
 *
 * 1. 一个 integer 包含值 123
 * 2. 一个包含两个元素的嵌套列表：
 * ⁠   i.  一个 integer 包含值 456
 * ⁠   ii. 一个包含一个元素的嵌套列表
 * ⁠        a. 一个 integer 包含值 789
 *
 *
 *
 *
 */
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
func deserialize(s string) *NestedInteger {
	stack := NewStack()
	from := 0
	for from < len(s) {
		if s[from] == '[' {
			n := new(NestedInteger)
			stack.Push(n)
			from++
			continue
		}
		if s[from] == ',' {
			from++
			continue
		}
		if s[from] == '-' || (s[from] >= '0' && s[from] <= '9') {
			val, to := parseInt(s, from)
			from = to
			if stack.Size() == 0 {
				n := new(NestedInteger)
				n.SetInteger(val)
				stack.Push(n)
			} else {
				stack.Top().SetInteger(val)
			}
			continue
		}
		if s[from] == ']' {
			n := stack.Pop()
			if stack.Size() == 0 {
				stack.Push(n)
			} else {
				stack.Top().Add(*n)
			}
			from++
		}
	}
	return stack.Top()
}

func parseInt(s string, from int) (int, int) {
	val := 0
	isNegative := false
	if s[from] == '-' {
		isNegative = true
		from++
	}
	for from < len(s) && s[from] >= '0' && s[from] <= '9' {
		val = val*10 + int(s[from]-'0')
		from++
	}
	if isNegative {
		val = val * -1
	}
	return val, from

}

type Stack struct {
	elements []*NestedInteger
}

func NewStack() *Stack {
	return &Stack{
		elements: make([]*NestedInteger, 0),
	}
}

func (s *Stack) Size() int {
	return len(s.elements)
}

func (s *Stack) Push(node *NestedInteger) {
	s.elements = append(s.elements, node)
}

func (s *Stack) Pop() *NestedInteger {
	size := s.Size()
	val := s.elements[size-1]
	s.elements = s.elements[:size-1]
	return val
}

func (s *Stack) Top() *NestedInteger {
	size := s.Size()
	return s.elements[size-1]
}

