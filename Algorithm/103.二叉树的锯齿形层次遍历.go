/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层次遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/description/
 *
 * algorithms
 * Medium (49.27%)
 * Likes:    60
 * Dislikes: 0
 * Total Accepted:    10.6K
 * Total Submissions: 21.6K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
 *
 * 例如：
 * 给定二叉树 [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 * 返回锯齿形层次遍历如下：
 *
 * [
 * ⁠ [3],
 * ⁠ [20,9],
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
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	isLeftToRight := true
	for len(queue) > 0 {
		newQueue := make([]*TreeNode, 0)
		values := make([]int, len(queue))
		for i := len(queue) - 1; i >= 0; i-- {
			values[len(queue)-1-i] = queue[i].Val
			if isLeftToRight {
				if queue[i].Left != nil {
					newQueue = append(newQueue, queue[i].Left)
				}
				if queue[i].Right != nil {
					newQueue = append(newQueue, queue[i].Right)
				}
			} else {
				if queue[i].Right != nil {
					newQueue = append(newQueue, queue[i].Right)
				}
				if queue[i].Left != nil {
					newQueue = append(newQueue, queue[i].Left)
				}
			}
		}
		isLeftToRight = !isLeftToRight
		queue = newQueue
		result = append(result, values)
	}
	return result
}

