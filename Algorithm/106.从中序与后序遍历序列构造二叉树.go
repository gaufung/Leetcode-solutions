/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (60.64%)
 * Likes:    78
 * Dislikes: 0
 * Total Accepted:    7.8K
 * Total Submissions: 12.8K
 * Testcase Example:  '[9,3,15,20,7]\n[9,15,7,20,3]'
 *
 * 根据一棵树的中序遍历与后序遍历构造二叉树。
 *
 * 注意:
 * 你可以假设树中没有重复的元素。
 *
 * 例如，给出
 *
 * 中序遍历 inorder = [9,3,15,20,7]
 * 后序遍历 postorder = [9,15,7,20,3]
 *
 * 返回如下的二叉树：
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	if len(inorder) == 1 {
		return &TreeNode{
			Val: inorder[0],
		}
	}
	n := len(postorder)
	rootValue := postorder[n-1]
	inIndex := findIndex(inorder, rootValue)
	leftInOrder := inorder[:inIndex]
	rightInOrder := inorder[inIndex+1:]
	leftPostorder := postorder[:inIndex]
	rightPostOrder := postorder[inIndex : n-1]
	root := &TreeNode{
		Val: rootValue,
	}
	root.Left = buildTree(leftInOrder, leftPostorder)
	root.Right = buildTree(rightInOrder, rightPostOrder)
	return root
}

func removeKth(nums []int, k int) []int {
	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if i != k {
			result = append(result, nums[i])
		}
	}
	return result
}

func findIndex(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	return -1
}

