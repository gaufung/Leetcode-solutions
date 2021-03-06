/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
 *
 * https://leetcode-cn.com/problems/path-sum-ii/description/
 *
 * algorithms
 * Medium (55.04%)
 * Likes:    84
 * Dislikes: 0
 * Total Accepted:    8.9K
 * Total Submissions: 16.2K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,5,1]\n22'
 *
 * 给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例:
 * 给定如下二叉树，以及目标和 sum = 22，
 *
 * ⁠             5
 * ⁠            / \
 * ⁠           4   8
 * ⁠          /   / \
 * ⁠         11  13  4
 * ⁠        /  \    / \
 * ⁠       7    2  5   1
 *
 *
 * 返回:
 *
 * [
 * ⁠  [5,4,11,2],
 * ⁠  [5,8,4,5]
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

