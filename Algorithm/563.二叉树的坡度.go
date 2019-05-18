/*
 * @lc app=leetcode.cn id=563 lang=golang
 *
 * [563] 二叉树的坡度
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
