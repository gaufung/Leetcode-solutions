/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 *
 * https://leetcode-cn.com/problems/reverse-linked-list-ii/description/
 *
 * algorithms
 * Medium (43.76%)
 * Likes:    137
 * Dislikes: 0
 * Total Accepted:    10.4K
 * Total Submissions: 23.8K
 * Testcase Example:  '[1,2,3,4,5]\n2\n4'
 *
 * 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
 *
 * 说明:
 * 1 ≤ m ≤ n ≤ 链表长度。
 *
 * 示例:
 *
 * 输入: 1->2->3->4->5->NULL, m = 2, n = 4
 * 输出: 1->4->3->2->5->NULL
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	M := m
	first := new(ListNode)
	first.Next = head
	second := new(ListNode)
	second.Next = head
	for m > 1 {
		first = first.Next
		second = second.Next
		m--
		n--
	}
	second = second.Next
	for n > 1 {
		temp := second.Next
		second.Next = temp.Next
		temp.Next = first.Next
		first.Next = temp
		n--
	}
	if M == 1 {
		return first.Next
	} else {
		return head
	}

}

