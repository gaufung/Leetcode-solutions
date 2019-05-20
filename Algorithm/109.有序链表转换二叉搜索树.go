/*
 * @lc app=leetcode.cn id=109 lang=golang
 *
 * [109] 有序链表转换二叉搜索树
 *
 * https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/description/
 *
 * algorithms
 * Medium (65.73%)
 * Likes:    66
 * Dislikes: 0
 * Total Accepted:    6K
 * Total Submissions: 9.1K
 * Testcase Example:  '[-10,-3,0,5,9]'
 *
 * 给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。
 *
 * 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
 *
 * 示例:
 *
 * 给定的有序链表： [-10, -3, 0, 5, 9],
 *
 * 一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：
 *
 * ⁠     0
 * ⁠    / \
 * ⁠  -3   9
 * ⁠  /   /
 * ⁠-10  5
 *
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{
			Val: head.Val,
		}
	}
	if head.Next.Next == nil {
		first := &TreeNode{Val: head.Val}
		second := &TreeNode{Val: head.Next.Val}
		second.Left = first
		return second
	}
	if head.Next.Next.Next == nil {
		first := &TreeNode{Val: head.Val}
		second := &TreeNode{Val: head.Next.Val}
		third := &TreeNode{Val: head.Next.Next.Val}
		second.Left = first
		second.Right = third
	}
	fast, slow := new(ListNode), new(ListNode)
	fast = head
	slow.Next = head
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next
		}
	}
	root := &TreeNode{Val: slow.Next.Val}
	right := sortedListToBST(slow.Next.Next)
	slow.Next = nil
	left := sortedListToBST(head)
	root.Left = left
	root.Right = right
	return root

}

