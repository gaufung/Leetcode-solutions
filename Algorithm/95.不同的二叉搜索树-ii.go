/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	if n <= 0 {
		return []*TreeNode{}
	}
	return generate(1, n)
}
func generate(from, to int) []*TreeNode {
	if from > to {
		return []*TreeNode{
			nil,
		}
	}
	if from == to {
		return []*TreeNode{
			&TreeNode{Val: from},
		}
	}
	result := make([]*TreeNode, 0)
	for i := from; i <= to; i++ {
		var leftTree []*TreeNode
		var rightTree []*TreeNode
		if i == from {
			leftTree = []*TreeNode{nil}
		} else {
			leftTree = generate(from, i-1)
		}
		if i == to {
			rightTree = []*TreeNode{nil}
		} else {
			rightTree = generate(i+1, to)
		}
		for _, left := range leftTree {
			for _, right := range rightTree {
				node := &TreeNode{Val: i}
				node.Left = left
				node.Right = right
				result = append(result, node)
			}
		}
	}
	return result
}
