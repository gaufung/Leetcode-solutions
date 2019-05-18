/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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

