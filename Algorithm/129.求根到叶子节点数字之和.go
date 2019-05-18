/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根到叶子节点数字之和
 *
 * https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/description/
 *
 * algorithms
 * Medium (56.98%)
 * Likes:    57
 * Dislikes: 0
 * Total Accepted:    5.9K
 * Total Submissions: 10.3K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个二叉树，它的每个结点都存放一个 0-9 的数字，每条从根到叶子节点的路径都代表一个数字。
 *
 * 例如，从根到叶子节点路径 1->2->3 代表数字 123。
 *
 * 计算从根到叶子节点生成的所有数字之和。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例 1:
 *
 * 输入: [1,2,3]
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   3
 * 输出: 25
 * 解释:
 * 从根到叶子节点路径 1->2 代表数字 12.
 * 从根到叶子节点路径 1->3 代表数字 13.
 * 因此，数字总和 = 12 + 13 = 25.
 *
 * 示例 2:
 *
 * 输入: [4,9,0,5,1]
 * ⁠   4
 * ⁠  / \
 * ⁠ 9   0
 * / \
 * 5   1
 * 输出: 1026
 * 解释:
 * 从根到叶子节点路径 4->9->5 代表数字 495.
 * 从根到叶子节点路径 4->9->1 代表数字 491.
 * 从根到叶子节点路径 4->0 代表数字 40.
 * 因此，数字总和 = 495 + 491 + 40 = 1026.
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
