/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
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

