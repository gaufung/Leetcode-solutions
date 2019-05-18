/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	s := NewStack()
	goAlongLeftBranch(root, s)
	if s.Size() == 0 {
		return true
	}
	firstNode := s.Pop()
	if firstNode.Right != nil {
		goAlongLeftBranch(firstNode.Right, s)
	}
	minValue := firstNode.Val
	return valid(s, minValue)

}

func valid(s *Stack, minValue int) bool {
	for s.Size() > 0 {
		node := s.Pop()
		if node.Val <= minValue {
			return false
		} else {
			minValue = node.Val
		}
		goAlongLeftBranch(node.Right, s)
	}
	return true
}

func goAlongLeftBranch(node *TreeNode, s *Stack) {
	for node != nil {
		s.Push(node)
		node = node.Left
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
