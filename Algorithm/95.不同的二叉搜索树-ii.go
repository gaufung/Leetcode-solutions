/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
 *
 * https://leetcode-cn.com/problems/unique-binary-search-trees-ii/description/
 *
 * algorithms
 * Medium (54.72%)
 * Likes:    131
 * Dislikes: 0
 * Total Accepted:    5.4K
 * Total Submissions: 9.9K
 * Testcase Example:  '3'
 *
 * 给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。
 *
 * 示例:
 *
 * 输入: 3
 * 输出:
 * [
 * [1,null,3,2],
 * [3,2,null,1],
 * [3,1,null,null,2],
 * [2,1,3],
 * [1,null,2,null,3]
 * ]
 * 解释:
 * 以上的输出对应以下 5 种不同结构的二叉搜索树：
 *
 * ⁠  1         3     3      2      1
 * ⁠   \       /     /      / \      \
 * ⁠    3     2     1      1   3      2
 * ⁠   /     /       \                 \
 * ⁠  2     1         2                 3
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

