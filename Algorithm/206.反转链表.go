/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 *
 * https://leetcode-cn.com/problems/reverse-linked-list/description/
 *
 * algorithms
 * Easy (61.22%)
 * Likes:    426
 * Dislikes: 0
 * Total Accepted:    57.7K
 * Total Submissions: 94.2K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 反转一个单链表。
 *
 * 示例:
 *
 * 输入: 1->2->3->4->5->NULL
 * 输出: 5->4->3->2->1->NULL
 *
 * 进阶:
 * 你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	h, _ := reverse(head)
	return h
}

func reverse(node *ListNode) (*ListNode, *ListNode) {
	if node.Next == nil {
		return node, node
	}
	n1, n2 := reverse(node.Next)
	node.Next = nil
	n2.Next = node
	return n1, node
}

