/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	count := 0
	pos := new(ListNode)
	pos.Next = head
	for pos.Next != nil {
		count++
		pos = pos.Next
	}
	pos.Next = head
	k = count - k%count
	cur := new(ListNode)
	cur.Next = head
	for k > 0 {
		k--
		cur = cur.Next
	}
	result := cur.Next
	cur.Next = nil
	return result
}

