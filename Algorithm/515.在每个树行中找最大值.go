/*
 * @lc app=leetcode.cn id=515 lang=golang
 *
 * [515] 在每个树行中找最大值
 *
 * https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/description/
 *
 * algorithms
 * Medium (56.35%)
 * Likes:    36
 * Dislikes: 0
 * Total Accepted:    2.7K
 * Total Submissions: 4.7K
 * Testcase Example:  '[1,3,2,5,3,null,9]'
 *
 * 您需要在二叉树的每一行中找到最大的值。
 *
 * 示例：
 *
 *
 * 输入:
 *
 * ⁠         1
 * ⁠        / \
 * ⁠       3   2
 * ⁠      / \   \
 * ⁠     5   3   9
 *
 * 输出: [1, 3, 9]
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
func largestValues(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		newQueue := make([]*TreeNode, 0)
		val := queue[0].Val
		for _, q := range queue {
			val = max(val, q.Val)
			if q.Left != nil {
				newQueue = append(newQueue, q.Left)
			}
			if q.Right != nil {
				newQueue = append(newQueue, q.Right)
			}
		}
		result = append(result, val)
		queue = newQueue
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
