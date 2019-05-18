/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSame(root.Left, root.Right)
}

func isSame(p, q *TreeNode) bool {
	if q == nil && p == nil {
		return true
	}
	if q != nil && p == nil {
		return false
	}
	if q == nil && p != nil {
		return false
	}
	if q.Val != p.Val {
		return false
	}
	return isSame(p.Left, q.Right) && isSame(p.Right, q.Left)
}

