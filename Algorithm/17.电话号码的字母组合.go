/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 *
 * https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (49.09%)
 * Likes:    318
 * Dislikes: 0
 * Total Accepted:    24.5K
 * Total Submissions: 49.8K
 * Testcase Example:  '"23"'
 *
 * 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
 *
 * 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
 *
 *
 *
 * 示例:
 *
 * 输入："23"
 * 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 *
 *
 * 说明:
 * 尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
 *
 */
func letterCombinations(digits string) []string {
	keyboard := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	result := make([]string, 0)
	if len(digits) == 0 {
		return result
	}
	stack := NewStack()
	stack.Push(-1)
	for stack.Size() > 0 {
		index := stack.Top()
		index += 1
		if index < len(keyboard[digits[stack.Size()-1]]) {
			stack.Pop()
			stack.Push(index)
		} else {
			stack.Pop()
			continue
		}
		for i := stack.Size(); i < len(digits); i++ {
			stack.Push(0)
		}
		allIndexes := stack.Values()
		items := make([]byte, len(digits))
		for i := 0; i < len(digits); i++ {
			digit := digits[i]
			index := allIndexes[i]
			items[i] = keyboard[digit][index]
		}
		result = append(result, string(items))
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

func (s *Stack) Pop() {
	s.elements = s.elements[:len(s.elements)-1]
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

