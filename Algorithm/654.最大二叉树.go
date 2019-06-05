/*
 * @lc app=leetcode.cn id=654 lang=golang
 *
 * [654] 最大二叉树
 *
 * https://leetcode-cn.com/problems/maximum-binary-tree/description/
 *
 * algorithms
 * Medium (74.22%)
 * Likes:    57
 * Dislikes: 0
 * Total Accepted:    4.1K
 * Total Submissions: 5.5K
 * Testcase Example:  '[3,2,1,6,0,5]'
 *
 * 给定一个不含重复元素的整数数组。一个以此数组构建的最大二叉树定义如下：
 *
 *
 * 二叉树的根是数组中的最大元素。
 * 左子树是通过数组中最大值左边部分构造出的最大二叉树。
 * 右子树是通过数组中最大值右边部分构造出的最大二叉树。
 *
 *
 * 通过给定的数组构建最大二叉树，并且输出这个树的根节点。
 *
 * Example 1:
 *
 *
 * 输入: [3,2,1,6,0,5]
 * 输入: 返回下面这棵树的根节点：
 *
 * ⁠     6
 * ⁠   /   \
 * ⁠  3     5
 * ⁠   \    /
 * ⁠    2  0
 * ⁠      \
 * ⁠       1
 *
 *
 * 注意:
 *
 *
 * 给定的数组的大小在 [1, 1000] 之间。
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
func constructMaximumBinaryTree(nums []int) *TreeNode {
	n := len(nums)
	if n <= 0 {
		return nil
	}
	index := findMaxIndex(nums)
	root := &TreeNode{
		Val: nums[index],
	}
	root.Left = constructMaximumBinaryTree(nums[:index])
	root.Right = constructMaximumBinaryTree(nums[index+1:])
	return root

}

func findMaxIndex(nums []int) int {
	index := 0
	val := nums[index]
	for i := 1; i < len(nums); i++ {
		if nums[i] > val {
			val = nums[i]
			index = i
		}
	}
	return index
}



