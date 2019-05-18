/*
 * @lc app=leetcode.cn id=235 lang=golang
 *
 * [235] 二叉搜索树的最近公共祖先
 */
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	sp := NewStack()
	PreOrder(root, p, sp)
	sq := NewStack()
	PreOrder(root, q, sq)
	pPaths := sp.Values()
	qPaths := sq.Values()
	i, j := 0, 0
	for i < len(pPaths) && j < len(qPaths) {
		if pPaths[i] == qPaths[j] {
			i++
			j++
		} else {
			break
		}
	}
	return pPaths[i-1]

}

func PreOrder(node, expect *TreeNode, s *Stack) {
	if node != nil {
		s.Push(node)
		PreOrder(node.Left, expect, s)
		PreOrder(node.Right, expect, s)
		if s.Size() > 0 && s.Top() != expect {
			s.Pop()
		}

	}
}

type Stack struct {
	elements []*TreeNode
}

func NewStack() *Stack {
	return &Stack{
		elements: make([]*TreeNode, 0),
	}
}

func (s *Stack) Size() int {
	return len(s.elements)
}

func (s *Stack) Push(node *TreeNode) {
	s.elements = append(s.elements, node)
}

func (s *Stack) Pop() *TreeNode {
	size := s.Size()
	val := s.elements[size-1]
	s.elements = s.elements[:size-1]
	return val
}

func (s *Stack) Top() *TreeNode {
	size := s.Size()
	return s.elements[size-1]
}

func (s *Stack) Values() []*TreeNode {
	return s.elements
}