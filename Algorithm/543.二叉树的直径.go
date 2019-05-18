/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	length := height(root.Left) + height(root.Right)
	left := diameterOfBinaryTree(root.Left)
	right := diameterOfBinaryTree(root.Right)
	return max(length, max(left, right))
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
