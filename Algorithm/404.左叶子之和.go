/*
 * @lc app=leetcode.cn id=404 lang=golang
 *
 * [404] 左叶子之和
 *
 * https://leetcode-cn.com/problems/sum-of-left-leaves/description/
 *
 * algorithms
 * Easy (49.93%)
 * Likes:    77
 * Dislikes: 0
 * Total Accepted:    7.1K
 * Total Submissions: 14.2K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 计算给定二叉树的所有左叶子之和。
 *
 * 示例：
 *
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * 在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24
 *
 *
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
func sumOfLeftLeaves(root *TreeNode) int {
	return sum(root, false)
	// if root == nil || (root.Left == nil && root.Right == nil) {
	// 	return 0
	// }
	// if root.Left != nil && root.Right == nil {
	// 	return root.Left.Val + sumOfLeftLeaves(root.Left)
	// }
	// if root.Right != nil && root.Left == nil {
	// 	return sumOfLeftLeaves(root.Right)
	// }
	// return root.Left.Val + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}

func sum(node *TreeNode, isLeft bool) int {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil {
		if isLeft {
			return node.Val
		}
		return 0
	}
	return sum(node.Left, true) + sum(node.Right, false)
}

