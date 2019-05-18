/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			return [][]int{
				[]int{sum},
			}
		} else {
			return [][]int{}
		}
	}
	if root.Left == nil && root.Right != nil {
		rightPaths := pathSum(root.Right, sum-root.Val)
		result := make([][]int, 0)
		for _, path := range rightPaths {
			result = append(result, append([]int{root.Val}, path...))
		}
		return result
	}
	if root.Left != nil && root.Right == nil {
		leftPaths := pathSum(root.Left, sum-root.Val)
		result := make([][]int, 0)
		for _, path := range leftPaths {
			result = append(result, append([]int{root.Val}, path...))
		}
		return result
	}
	result := make([][]int, 0)
	leftPaths := pathSum(root.Left, sum-root.Val)
	for _, path := range leftPaths {
		result = append(result, append([]int{root.Val}, path...))
	}
	rightPaths := pathSum(root.Right, sum-root.Val)
	for _, path := range rightPaths {
		result = append(result, append([]int{root.Val}, path...))
	}
	return result
}

