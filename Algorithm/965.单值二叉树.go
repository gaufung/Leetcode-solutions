/*
 * @lc app=leetcode.cn id=965 lang=golang
 *
 * [965] 单值二叉树
 *
 * https://leetcode-cn.com/problems/univalued-binary-tree/description/
 *
 * algorithms
 * Easy (65.00%)
 * Likes:    20
 * Dislikes: 0
 * Total Accepted:    5.6K
 * Total Submissions: 8.7K
 * Testcase Example:  '[1,1,1,1,1,null,1]'
 *
 * 如果二叉树每个节点都具有相同的值，那么该二叉树就是单值二叉树。
 *
 * 只有给定的树是单值二叉树时，才返回 true；否则返回 false。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：[1,1,1,1,1,null,1]
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 *
 * 输入：[2,2,2,5,2]
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 给定树的节点数范围是 [1, 100]。
 * 每个节点的值都是整数，范围为 [0, 99] 。
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
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left == nil && root.Right != nil {
		if root.Val != root.Right.Val {
			return false
		}
		return isUnivalTree(root.Right)
	}
	if root.Left != nil && root.Right == nil {
		if root.Val != root.Left.Val {
			return false
		}
		return isUnivalTree(root.Left)
	} else {
		if root.Val != root.Left.Val || root.Val != root.Right.Val {
			return false
		}
		return isUnivalTree(root.Left) && isUnivalTree(root.Right)
	}
}

