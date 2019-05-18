/*
 * @lc app=leetcode.cn id=450 lang=golang
 *
 * [450] 删除二叉搜索树中的节点
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			successor := lowestLeftNode(root.Right)
			successor.Right = deleteMin(root.Right)
			successor.Left = root.Left
			return successor
		}
	}
}

func lowestLeftNode(node *TreeNode) *TreeNode {
	if node.Left == nil {
		return node
	}
	return lowestLeftNode(node.Left)
}

func deleteMin(node *TreeNode) *TreeNode {
	if node.Left == nil {
		return node.Right
	}
	node.Left = deleteMin(node.Left)
	return node
}

