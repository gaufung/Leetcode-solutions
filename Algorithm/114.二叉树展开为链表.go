/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flat(root)
}

func flat(node *TreeNode) (*TreeNode, *TreeNode) {
	if node.Left == nil && node.Right == nil {
		return node, node
	}
	left, right := node.Left, node.Right
	last := node
	if left != nil {
		head, tail := flat(left)
		last.Right = head
		last = tail
	}
	if right != nil {
		head, tail := flat(right)
		last.Right = head
		last = tail
	}
	node.Left = nil
	return node, last
}

