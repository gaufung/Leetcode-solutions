/*
 * @lc app=leetcode.cn id=563 lang=golang
 *
 * [563] 二叉树的坡度
 *
 * https://leetcode-cn.com/problems/binary-tree-tilt/description/
 *
 * algorithms
 * Easy (47.94%)
 * Likes:    30
 * Dislikes: 0
 * Total Accepted:    3.6K
 * Total Submissions: 7.4K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个二叉树，计算整个树的坡度。
 *
 * 一个树的节点的坡度定义即为，该节点左子树的结点之和和右子树结点之和的差的绝对值。空结点的的坡度是0。
 *
 * 整个树的坡度就是其所有节点的坡度之和。
 *
 * 示例:
 *
 *
 * 输入:
 * ⁠        1
 * ⁠      /   \
 * ⁠     2     3
 * 输出: 1
 * 解释:
 * 结点的坡度 2 : 0
 * 结点的坡度 3 : 0
 * 结点的坡度 1 : |2-3| = 1
 * 树的坡度 : 0 + 0 + 1 = 1
 *
 *
 * 注意:
 *
 *
 * 任何子树的结点的和不会超过32位整数的范围。
 * 坡度的值不会超过32位整数的范围。
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
func findTilt(root *TreeNode) int {
	_, tile := SumAnTilt(root)
	return tile
}

func SumAnTilt(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}
	sumLeft, tiltLeft := SumAnTilt(node.Left)
	sumRight, tiltRight := SumAnTilt(node.Right)
	sum := sumLeft + sumRight + node.Val
	tile := tiltLeft + tiltRight + abs(sumLeft, sumRight)
	return sum, tile
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

