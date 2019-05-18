/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
 *
 * https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/description/
 *
 * algorithms
 * Medium (60.67%)
 * Likes:    108
 * Dislikes: 0
 * Total Accepted:    6.7K
 * Total Submissions: 11.1K
 * Testcase Example:  '[1,2,5,3,4,null,6]'
 *
 * 给定一个二叉树，原地将它展开为链表。
 *
 * 例如，给定二叉树
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   5
 * ⁠/ \   \
 * 3   4   6
 *
 * 将其展开为：
 *
 * 1
 * ⁠\
 * ⁠ 2
 * ⁠  \
 * ⁠   3
 * ⁠    \
 * ⁠     4
 * ⁠      \
 * ⁠       5
 * ⁠        \
 * ⁠         6
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
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flat(root)
}

func flat(node *TreeNode) (*TreeNode, *TreeNode) {
	if node.Left == nil && node.Right == nil {
		return node, node
	}
	left, right := node.Left, node.Right
	last := node
	if left != nil {
		head, tail := flat(left)
		last.Right = head
		last = tail
	}
	if right != nil {
		head, tail := flat(right)
		last.Right = head
		last = tail
	}
	node.Left = nil
	return node, last
}
