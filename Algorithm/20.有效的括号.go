/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 *
 * https://leetcode-cn.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (38.07%)
 * Likes:    804
 * Dislikes: 0
 * Total Accepted:    74.8K
 * Total Submissions: 196.5K
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
 *
 * 有效字符串需满足：
 *
 *
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 *
 *
 * 注意空字符串可被认为是有效字符串。
 *
 * 示例 1:
 *
 * 输入: "()"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: "()[]{}"
 * 输出: true
 *
 *
 * 示例 3:
 *
 * 输入: "(]"
 * 输出: false
 *
 *
 * 示例 4:
 *
 * 输入: "([)]"
 * 输出: false
 *
 *
 * 示例 5:
 *
 * 输入: "{[]}"
 * 输出: true
 *
 */
func isValid(s string) bool {
	ss := newStack()
	for _, ch := range s {
		if ss.isEmpty() {
			ss.push(ch)
		} else {
			if isMatch(ss.peek(), ch) {
				ss.pop()
			} else {
				ss.push(ch)
			}
		}
	}
	return ss.isEmpty()
}

func isMatch(left, right rune) bool {
	return (left == '(' && right == ')') || (left == '{' && right == '}') || (left == '[' && right == ']')
}

type stack struct {
	values []rune
	pos    int
}

func newStack() *stack {
	return &stack{make([]rune, 1), -1}
}

func (s *stack) pop() rune {
	val := s.values[s.pos]
	s.pos--
	return val
}

func (s *stack) push(val rune) {
	if s.pos == len(s.values)-1 {
		backup := make([]rune, 2*len(s.values))
		copy(backup, s.values)
		s.values = backup
	}
	s.pos++
	s.values[s.pos] = val
}

func (s *stack) isEmpty() bool {
	return s.pos == -1
}
func (s *stack) peek() rune {
	return s.values[s.pos]
}

