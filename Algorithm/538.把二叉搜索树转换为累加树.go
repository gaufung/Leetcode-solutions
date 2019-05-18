/*
 * @lc app=leetcode.cn id=538 lang=golang
 *
 * [538] 把二叉搜索树转换为累加树
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	a := 0
	inorder(root, &a)
	return root
}

func inorder(node *TreeNode, val *int) {
	if node != nil {
		inorder(node.Right, val)
		node.Val += *val
		*val = node.Val
		inorder(node.Left, val)
	}
}

