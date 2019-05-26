import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
 *
 * https://leetcode-cn.com/problems/binary-tree-paths/description/
 *
 * algorithms
 * Easy (58.25%)
 * Likes:    123
 * Dislikes: 0
 * Total Accepted:    9.4K
 * Total Submissions: 16.2K
 * Testcase Example:  '[1,2,3,null,5]'
 *
 * 给定一个二叉树，返回所有从根节点到叶子节点的路径。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例:
 *
 * 输入:
 *
 * ⁠  1
 * ⁠/   \
 * 2     3
 * ⁠\
 * ⁠ 5
 *
 * 输出: ["1->2->5", "1->3"]
 *
 * 解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
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
func binaryTreePaths(root *TreeNode) []string {
	paths := binaryPath(root)
	result := make([]string, 0)
	for _, path := range paths {
		result = append(result, strings.Join(path, "->"))
	}
	return result
}

func binaryPath(node *TreeNode) [][]string {
	if node == nil {
		return [][]string{}
	}
	if node.Left == nil && node.Right == nil {
		return [][]string{
			[]string{strconv.Itoa(node.Val)},
		}
	}
	if node.Left == nil && node.Right != nil {
		result := make([][]string, 0)
		rightResult := binaryPath(node.Right)
		for _, right := range rightResult {
			result = append(result, append([]string{strconv.Itoa(node.Val)}, right...))
		}
		return result
	}
	if node.Left != nil && node.Right == nil {
		result := make([][]string, 0)
		leftResult := binaryPath(node.Left)
		for _, left := range leftResult {
			result = append(result, append([]string{strconv.Itoa(node.Val)}, left...))
		}
		return result
	} else {
		result := make([][]string, 0)
		leftResult := binaryPath(node.Left)
		for _, left := range leftResult {
			result = append(result, append([]string{strconv.Itoa(node.Val)}, left...))
		}
		rightResult := binaryPath(node.Right)
		for _, right := range rightResult {
			result = append(result, append([]string{strconv.Itoa(node.Val)}, right...))
		}
		return result
	}
}

