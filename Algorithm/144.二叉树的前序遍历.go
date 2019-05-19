/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-preorder-traversal/description/
 *
 * algorithms
 * Medium (59.52%)
 * Likes:    109
 * Dislikes: 0
 * Total Accepted:    22.2K
 * Total Submissions: 37.2K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的 前序 遍历。
 *
 * 示例:
 *
 * 输入: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 *
 * 输出: [1,2,3]
 *
 *
 * 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
 *
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	s := NewStack()
	result := visitAlongLeftBranch(root, s)
	for s.Size() > 0 {
		node := s.Pop()
		items := visitAlongLeftBranch(node, s)
		result = append(result, items...)
	}
	return result
}

func visitAlongLeftBranch(node *TreeNode, s *Stack) []int {
	result := make([]int, 0)
	for node != nil {
		result = append(result, node.Val)
		if node.Right != nil {
			s.Push(node.Right)
		}
		node = node.Left
	}
	return result
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

