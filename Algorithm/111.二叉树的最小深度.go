/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left == nil && root.Right != nil {
		return 1 + minDepth(root.Right)
	}
	if root.Right == nil && root.Left != nil {
		return 1 + minDepth(root.Left)
	}
	return 1 + min(minDepth(root.Left), minDepth(root.Right))

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

