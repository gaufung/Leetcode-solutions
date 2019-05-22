import "strings"

/*
 * @lc app=leetcode.cn id=331 lang=golang
 *
 * [331] 验证二叉树的前序序列化
 *
 * https://leetcode-cn.com/problems/verify-preorder-serialization-of-a-binary-tree/description/
 *
 * algorithms
 * Medium (41.66%)
 * Likes:    30
 * Dislikes: 0
 * Total Accepted:    1.5K
 * Total Submissions: 3.7K
 * Testcase Example:  '"9,3,4,#,#,1,#,#,2,#,6,#,#"'
 *
 * 序列化二叉树的一种方法是使用前序遍历。当我们遇到一个非空节点时，我们可以记录下这个节点的值。如果它是一个空节点，我们可以使用一个标记值记录，例如 #。
 *
 * ⁠    _9_
 * ⁠   /   \
 * ⁠  3     2
 * ⁠ / \   / \
 * ⁠4   1  #  6
 * / \ / \   / \
 * # # # #   # #
 *
 *
 * 例如，上面的二叉树可以被序列化为字符串 "9,3,4,#,#,1,#,#,2,#,6,#,#"，其中 # 代表一个空节点。
 *
 * 给定一串以逗号分隔的序列，验证它是否是正确的二叉树的前序序列化。编写一个在不重构树的条件下的可行算法。
 *
 * 每个以逗号分隔的字符或为一个整数或为一个表示 null 指针的 '#' 。
 *
 * 你可以认为输入格式总是有效的，例如它永远不会包含两个连续的逗号，比如 "1,,3" 。
 *
 * 示例 1:
 *
 * 输入: "9,3,4,#,#,1,#,#,2,#,6,#,#"
 * 输出: true
 *
 * 示例 2:
 *
 * 输入: "1,#"
 * 输出: false
 *
 *
 * 示例 3:
 *
 * 输入: "9,#,#,1"
 * 输出: false
 *
 */
func isValidSerialization(preorder string) bool {
	nodes := strings.Split(preorder, ",")
	stack := NewStack()
	for _, node := range nodes {
		stack.Push(node)
		if node == "#" {
			check(stack)
		}
	}
	return stack.Size() == 1 && stack.Top() == "#"
}

func check(stack *Stack) {
	if stack.Size() > 2 {
		val1 := stack.Pop()
		val2 := stack.Pop()
		val3 := stack.Pop()
		if val1 == "#" && val2 == "#" && val3 != "#" {
			stack.Push("#")
			check(stack)
		} else {
			stack.Push(val3)
			stack.Push(val2)
			stack.Push(val1)
		}
	}
}

type Stack struct {
	elements []string
}

func NewStack() *Stack {
	return &Stack{
		elements: make([]string, 0),
	}
}

func (s *Stack) Size() int {
	return len(s.elements)
}

func (s *Stack) Push(node string) {
	s.elements = append(s.elements, node)
}

func (s *Stack) Pop() string {
	size := s.Size()
	val := s.elements[size-1]
	s.elements = s.elements[:size-1]
	return val
}

func (s *Stack) Top() string {
	size := s.Size()
	return s.elements[size-1]
}

// 9 3 4 # #
// 9 3 # 1 # #
// 9 3 # #
// 9 # 2 # 6
// 9 # 2 # #
// 9 # #
// # 
