/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 *
 * https://leetcode-cn.com/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (59.65%)
 * Likes:    206
 * Dislikes: 0
 * Total Accepted:    21.9K
 * Total Submissions: 36.7K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
 *
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 *
 *
 *
 * 示例:
 *
 * 给定 1->2->3->4, 你应该返回 2->1->4->3.
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
func swapPairs(head *ListNode) *ListNode {
	fast := head
	slow := new(ListNode)
	slow.Next = fast
	if fast == nil || fast.Next == nil {
		return head
	}
	result := fast.Next
	for fast != nil && fast.Next != nil {
		slow.Next = fast.Next
		fast.Next = fast.Next.Next
		slow.Next.Next = fast
		slow = fast
		fast = fast.Next
	}
	return result
}

