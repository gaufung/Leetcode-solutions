/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 *
 * https://leetcode-cn.com/problems/sort-list/description/
 *
 * algorithms
 * Medium (60.02%)
 * Likes:    189
 * Dislikes: 0
 * Total Accepted:    12.7K
 * Total Submissions: 21.2K
 * Testcase Example:  '[4,2,1,3]'
 *
 * 在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
 *
 * 示例 1:
 *
 * 输入: 4->2->1->3
 * 输出: 1->2->3->4
 *
 *
 * 示例 2:
 *
 * 输入: -1->5->3->4->0
 * 输出: -1->0->3->4->5
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	num := 0
	cur := head
	for cur != nil {
		num++
		cur = cur.Next
	}
	count := num / 2
	cur = head
	for count > 1 {
		cur = cur.Next
		count--
	}
	l1 := head
	l2 := cur.Next
	cur.Next = nil
	l1 = sortList(l1)
	l2 = sortList(l2)
	return mergeTwoLists(l1, l2)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	header := new(ListNode)
	cursor := header
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cursor.Next = l1
			l1 = l1.Next
		} else {
			cursor.Next = l2
			l2 = l2.Next
		}
		cursor = cursor.Next
	}
	if l1 != nil {
		cursor.Next = l1
	}
	if l2 != nil {
		cursor.Next = l2
	}
	return header.Next
}

