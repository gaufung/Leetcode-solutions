/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		result = append(result, queue[size-1].Val)
		newQueue := make([]*TreeNode, 0)
		for _, q := range queue {
			if q.Left != nil {
				newQueue = append(newQueue, q.Left)
			}
			if q.Right != nil {
				newQueue = append(newQueue, q.Right)
			}
			queue = newQueue
		}
	}
	return result
}

