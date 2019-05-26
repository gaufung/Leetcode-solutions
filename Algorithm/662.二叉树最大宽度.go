/*
 * @lc app=leetcode.cn id=662 lang=golang
 *
 * [662] 二叉树最大宽度
 *
 * https://leetcode-cn.com/problems/maximum-width-of-binary-tree/description/
 *
 * algorithms
 * Medium (32.45%)
 * Likes:    37
 * Dislikes: 0
 * Total Accepted:    1.5K
 * Total Submissions: 4.7K
 * Testcase Example:  '[1,3,2,5,3,null,9]'
 *
 * 给定一个二叉树，编写一个函数来获取这个树的最大宽度。树的宽度是所有层中的最大宽度。这个二叉树与满二叉树（full binary
 * tree）结构相同，但一些节点为空。
 *
 * 每一层的宽度被定义为两个端点（该层最左和最右的非空节点，两端点间的null节点也计入长度）之间的长度。
 *
 * 示例 1:
 *
 *
 * 输入:
 *
 * ⁠          1
 * ⁠        /   \
 * ⁠       3     2
 * ⁠      / \     \
 * ⁠     5   3     9
 *
 * 输出: 4
 * 解释: 最大值出现在树的第 3 层，宽度为 4 (5,3,null,9)。
 *
 *
 * 示例 2:
 *
 *
 * 输入:
 *
 * ⁠         1
 * ⁠        /
 * ⁠       3
 * ⁠      / \
 * ⁠     5   3
 *
 * 输出: 2
 * 解释: 最大值出现在树的第 3 层，宽度为 2 (5,3)。
 *
 *
 * 示例 3:
 *
 *
 * 输入:
 *
 * ⁠         1
 * ⁠        / \
 * ⁠       3   2
 * ⁠      /
 * ⁠     5
 *
 * 输出: 2
 * 解释: 最大值出现在树的第 2 层，宽度为 2 (3,2)。
 *
 *
 * 示例 4:
 *
 *
 * 输入:
 *
 * ⁠         1
 * ⁠        / \
 * ⁠       3   2
 * ⁠      /     \
 * ⁠     5       9
 * ⁠    /         \
 * ⁠   6           7
 * 输出: 8
 * 解释: 最大值出现在树的第 4 层，宽度为 8 (6,null,null,null,null,null,null,7)。
 *
 *
 * 注意: 答案在32位有符号整数的表示范围内。
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
func widthOfBinaryTree(root *TreeNode) int {
	res := 0
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		res = max(res, len(queue))
		newQueue := make([]*TreeNode, 0)
		for _, q := range queue {
			if q == nil {
				newQueue = append(newQueue, nil)
				newQueue = append(newQueue, nil)
			} else {
				newQueue = append(newQueue, []*TreeNode{q.Left, q.Right}...)
			}
		}
		newQueue = trimQueue(newQueue)
		queue = newQueue
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func trimQueue(queue []*TreeNode) []*TreeNode {
	i := 0
	for i < len(queue) && queue[i] == nil {
		i++
	}
	queue = queue[i:]
	j := len(queue) - 1
	for j >= 0 && queue[j] == nil {
		j--
	}
	return queue[:j+1]
}

