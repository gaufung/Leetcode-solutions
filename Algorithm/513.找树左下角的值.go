/*
 * @lc app=leetcode.cn id=513 lang=golang
 *
 * [513] 找树左下角的值
 *
 * https://leetcode-cn.com/problems/find-bottom-left-tree-value/description/
 *
 * algorithms
 * Medium (64.84%)
 * Likes:    36
 * Dislikes: 0
 * Total Accepted:    3.3K
 * Total Submissions: 5.1K
 * Testcase Example:  '[2,1,3]'
 *
 * 给定一个二叉树，在树的最后一行找到最左边的值。
 *
 * 示例 1:
 *
 *
 * 输入:
 *
 * ⁠   2
 * ⁠  / \
 * ⁠ 1   3
 *
 * 输出:
 * 1
 *
 *
 *
 *
 * 示例 2:
 *
 *
 * 输入:
 *
 * ⁠       1
 * ⁠      / \
 * ⁠     2   3
 * ⁠    /   / \
 * ⁠   4   5   6
 * ⁠      /
 * ⁠     7
 *
 * 输出:
 * 7
 *
 *
 *
 *
 * 注意: 您可以假设树（即给定的根节点）不为 NULL。
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

