/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	robbery, noroobery := 0, 0
	if root.Left == nil && root.Right == nil {
		robbery = root.Val
		noroobery = 0

	}
	if root.Left == nil && root.Right != nil {
		robbery = root.Val + rob(root.Right.Left) + rob(root.Right.Right)
		noroobery = rob(root.Right)
	}
	if root.Left != nil && root.Right == nil {
		robbery = root.Val + rob(root.Left.Left) + rob(root.Left.Right)
		noroobery = rob(root.Left)
	}
	if root.Left != nil && root.Right != nil {
		robbery = root.Val + rob(root.Left.Left) + rob(root.Left.Right) + rob(root.Right.Left) + rob(root.Right.Right)
		noroobery = rob(root.Right) + rob(root.Left)
	}
	return max(noroobery, robbery)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

