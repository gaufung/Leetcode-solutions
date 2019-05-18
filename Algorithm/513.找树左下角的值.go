/*
 * @lc app=leetcode.cn id=513 lang=golang
 *
 * [513] 找树左下角的值
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	if leftHeight >= rightHeight {
		return findBottomLeftValue(root.Left)
	} else {
		return findBottomLeftValue(root.Right)
	}
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

