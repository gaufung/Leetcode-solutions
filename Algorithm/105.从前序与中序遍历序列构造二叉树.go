/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
 *
 * algorithms
 * Medium (57.76%)
 * Likes:    156
 * Dislikes: 0
 * Total Accepted:    12.7K
 * Total Submissions: 22K
 * Testcase Example:  '[3,9,20,15,7]\n[9,3,15,20,7]'
 *
 * 根据一棵树的前序遍历与中序遍历构造二叉树。
 *
 * 注意:
 * 你可以假设树中没有重复的元素。
 *
 * 例如，给出
 *
 * 前序遍历 preorder = [3,9,20,15,7]
 * 中序遍历 inorder = [9,3,15,20,7]
 *
 * 返回如下的二叉树：
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	if len(preorder) == 1 && len(inorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	rootVal := preorder[0]
	j := 0
	for ; j < len(inorder); j++ {
		if inorder[j] == rootVal {
			break
		}
	}
	leftPreorder := preorder[1 : j+1]
	leftInorder := inorder[:j]
	rightPreorder := preorder[j+1:]
	rightInorder := inorder[j+1:]
	root := &TreeNode{Val: rootVal}
	root.Left = buildTree(leftPreorder, leftInorder)
	root.Right = buildTree(rightPreorder, rightInorder)
	return root

}

