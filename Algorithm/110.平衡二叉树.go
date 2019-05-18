/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	if leftHeight-rightHeight < -1 || leftHeight-rightHeight > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + max(height(node.Left), height(node.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
