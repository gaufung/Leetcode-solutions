/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 *
 * https://leetcode-cn.com/problems/reorder-list/description/
 *
 * algorithms
 * Medium (49.27%)
 * Likes:    72
 * Dislikes: 0
 * Total Accepted:    5K
 * Total Submissions: 10.1K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
 * 将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…
 *
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 *
 * 示例 1:
 *
 * 给定链表 1->2->3->4, 重新排列为 1->4->2->3.
 *
 * 示例 2:
 *
 * 给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	fast, slow := new(ListNode), new(ListNode)
	fast.Next = head
	slow.Next = head
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next
		}
	}
	head2, _ := reverse(slow.Next)
	slow.Next = nil
	head1 := head
	for head2 != nil {
		temp1 := head1.Next
		temp2 := head2.Next
		head1.Next = head2
		head2.Next = temp1
		head1 = temp1
		head2 = temp2
	}
}

func reverse(node *ListNode) (*ListNode, *ListNode) {
	if node.Next == nil {
		return node, node
	}
	head, tail := reverse(node.Next)
	node.Next = nil
	tail.Next = node
	return head, node
}

