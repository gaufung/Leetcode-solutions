/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根到叶子节点数字之和
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	paths := sumNode(root)
	result := 0
	for _, path := range paths {
		result += sum(path)
	}
	return result
}

func sum(nums []int) int {
	val := 0
	for _, num := range nums {
		val = val*10 + num
	}
	return val
}

func sumNode(node *TreeNode) [][]int {
	if node.Left == nil && node.Right == nil {
		return [][]int{
			[]int{node.Val},
		}
	}
	result := make([][]int, 0)
	if node.Left != nil {
		sums := sumNode(node.Left)
		for _, sum := range sums {
			result = append(result, append([]int{node.Val}, sum...))
		}
	}
	if node.Right != nil {
		sums := sumNode(node.Right)
		for _, sum := range sums {
			result = append(result, append([]int{node.Val}, sum...))
		}
	}
	return result
}

