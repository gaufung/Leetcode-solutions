/*
 * @lc app=leetcode.cn id=530 lang=golang
 *
 * [530] 二叉搜索树的最小绝对差
 *
 * https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/description/
 *
 * algorithms
 * Easy (53.25%)
 * Likes:    57
 * Dislikes: 0
 * Total Accepted:    4K
 * Total Submissions: 7.4K
 * Testcase Example:  '[1,null,3,2]'
 *
 * 给定一个所有节点为非负值的二叉搜索树，求树中任意两节点的差的绝对值的最小值。
 *
 * 示例 :
 *
 *
 * 输入:
 *
 * ⁠  1
 * ⁠   \
 * ⁠    3
 * ⁠   /
 * ⁠  2
 *
 * 输出:
 * 1
 *
 * 解释:
 * 最小绝对差为1，其中 2 和 1 的差的绝对值为 1（或者 2 和 3）。
 *
 *
 * 注意: 树中至少有2个节点。
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
func getMinimumDifference(root *TreeNode) int {
	val := make([]int, 0)
	inOrder(root, &val)
	res := val[len(val)-1] - val[0]
	for i := 1; i < len(val); i++ {
		if val[i]-val[i-1] < res {
			res = val[i] - val[i-1]
		}
	}
	return res
}

func inOrder(node *TreeNode, val *[]int) {
	if node != nil {
		inOrder(node.Left, val)
		*val = append(*val, node.Val)
		inOrder(node.Right, val)
	}
}

