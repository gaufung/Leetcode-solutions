/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
 *
 * https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/description/
 *
 * algorithms
 * Hard (35.03%)
 * Likes:    128
 * Dislikes: 0
 * Total Accepted:    6.9K
 * Total Submissions: 19.8K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个非空二叉树，返回其最大路径和。
 *
 * 本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
 *
 * 示例 1:
 *
 * 输入: [1,2,3]
 *
 * ⁠      1
 * ⁠     / \
 * ⁠    2   3
 *
 * 输出: 6
 *
 *
 * 示例 2:
 *
 * 输入: [-10,9,20,null,null,15,7]
 *
 * -10
 * / \
 * 9  20
 * /  \
 * 15   7
 *
 * 输出: 42
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
func maxPathSum(root *TreeNode) int {
	out, in := pathSum(root)
	return max(out, in)
}

func pathSum(node *TreeNode) (int, int) {
	if node == nil {
		return 0, min()
	}
	if node.Left == nil && node.Right == nil {
		return node.Val, node.Val
	} else {
		leftOut, leftIn := pathSum(node.Left)
		rightOut, rightIn := pathSum(node.Right)
		out := max(max(leftOut, rightOut)+node.Val, node.Val)
		in := max(max(max(max(rightOut+node.Val+leftOut, leftIn), rightIn), leftOut+node.Val), rightOut+node.Val)
		return out, in
	}
}

func min() int {
	a := ^uint(0)
	b := a >> 1
	return (int)(-b - 1)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

