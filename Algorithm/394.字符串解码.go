/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 *
 * https://leetcode-cn.com/problems/decode-string/description/
 *
 * algorithms
 * Medium (43.02%)
 * Likes:    77
 * Dislikes: 0
 * Total Accepted:    3.2K
 * Total Submissions: 7.4K
 * Testcase Example:  '"3[a]2[bc]"'
 *
 * 给定一个经过编码的字符串，返回它解码后的字符串。
 *
 * 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
 *
 * 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
 *
 * 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
 *
 * 示例:
 *
 *
 * s = "3[a]2[bc]", 返回 "aaabcbc".
 * s = "3[a2[c]]", 返回 "accaccacc".
 * s = "2[abc]3[cd]ef", 返回 "abcabccdcdcdef".
 *
 *
 */
func decodeString(s string) string {
	result := make([]byte, 0)
	n := len(s)
	if n <= 0 {
		return string(result)
	}
	stack := NewStack()
	index := 0
	for index < n {
		if isDigit(s[index]) {
			cnt, i := parseInt(s, index)
			r := NewRepeater(cnt)
			stack.Push(r)
			index = i
			continue
		}
		if s[index] == '[' {
			index++
			continue
		}
		if s[index] == ']' {
			r := stack.Pop()
			if stack.Size() > 0 {
				stack.Top().Add(r.Bytes())
			} else {
				result = append(result, r.Bytes()...)
			}
			index++
			continue
		}
		if stack.Size() > 0 {
			stack.Top().Add([]byte{s[index]})
		} else {
			result = append(result, s[index])
		}
		index++
	}
	return string(result)
}

func parseInt(s string, from int) (int, int) {
	val := 0
	for isDigit(s[from]) {
		val = val*10 + int(s[from]-'0')
		from++
	}
	return val, from
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

type Repeater struct {
	count int
	bytes []byte
}

func NewRepeater(cnt int) *Repeater {
	return &Repeater{
		count: cnt,
		bytes: make([]byte, 0),
	}
}

func (r *Repeater) Add(val []byte) {
	r.bytes = append(r.bytes, val...)
}

func (r *Repeater) Bytes() []byte {
	result := make([]byte, len(r.bytes)*r.count)
	for i := 0; i < r.count; i++ {
		for j := 0; j < len(r.bytes); j++ {
			result[i*len(r.bytes)+j] = r.bytes[j]
		}
	}
	return result
}

type Stack struct {
	elements []*Repeater
}

func NewStack() *Stack {
	return &Stack{
		elements: make([]*Repeater, 0),
	}
}

func (s *Stack) Size() int {
	return len(s.elements)
}

func (s *Stack) Push(node *Repeater) {
	s.elements = append(s.elements, node)
}

func (s *Stack) Pop() *Repeater {
	size := s.Size()
	val := s.elements[size-1]
	s.elements = s.elements[:size-1]
	return val
}

func (s *Stack) Top() *Repeater {
	size := s.Size()
	return s.elements[size-1]
}
