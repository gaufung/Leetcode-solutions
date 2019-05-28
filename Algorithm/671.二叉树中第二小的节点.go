/*
 * @lc app=leetcode.cn id=671 lang=golang
 *
 * [671] 二叉树中第二小的节点
 *
 * https://leetcode-cn.com/problems/second-minimum-node-in-a-binary-tree/description/
 *
 * algorithms
 * Easy (44.47%)
 * Likes:    39
 * Dislikes: 0
 * Total Accepted:    3.2K
 * Total Submissions: 7.2K
 * Testcase Example:  '[2,2,5,null,null,5,7]'
 *
 * 给定一个非空特殊的二叉树，每个节点都是正数，并且每个节点的子节点数量只能为 2 或
 * 0。如果一个节点有两个子节点的话，那么这个节点的值不大于它的子节点的值。
 *
 * 给出这样的一个二叉树，你需要输出所有节点中的第二小的值。如果第二小的值不存在的话，输出 -1 。
 *
 * 示例 1:
 *
 *
 * 输入:
 * ⁠   2
 * ⁠  / \
 * ⁠ 2   5
 * ⁠    / \
 * ⁠   5   7
 *
 * 输出: 5
 * 说明: 最小的值是 2 ，第二小的值是 5 。
 *
 *
 * 示例 2:
 *
 *
 * 输入:
 * ⁠   2
 * ⁠  / \
 * ⁠ 2   2
 *
 * 输出: -1
 * 说明: 最小的值是 2, 但是不存在第二小的值。
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
func findSecondMinimumValue(root *TreeNode) int {
	return findGreater(root, root.Val)
}

func findGreater(node *TreeNode, value int) int {
	if node == nil {
		return -1
	}
	if node.Val > value {
		return node.Val
	} else {
		left := findGreater(node.Left, value)
		right := findGreater(node.Right, value)
		if left == -1 && right == -1 {
			return -1
		} else if left == -1 && right != -1 {
			return right
		} else if left != -1 && right == -1 {
			return left
		} else {
			return min(left, right)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
