/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层次遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (55.86%)
 * Likes:    200
 * Dislikes: 0
 * Total Accepted:    25.7K
 * Total Submissions: 46K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。
 *
 * 例如:
 * 给定二叉树: [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 * 返回其层次遍历结果：
 *
 * [
 * ⁠ [3],
 * ⁠ [9,20],
 * ⁠ [15,7]
 * ]
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
func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		items := make([]int, 0)
		newQueue := make([]*TreeNode, 0)
		for _, q := range queue {
			if q != nil {
				items = append(items, q.Val)
				newQueue = append(newQueue, q.Left)
				newQueue = append(newQueue, q.Right)
			}
		}
		if len(items) > 0 {
			result = append(result, items)
		}
		queue = newQueue
	}
	return result
}

