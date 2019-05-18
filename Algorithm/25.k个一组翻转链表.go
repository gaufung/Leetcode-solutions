/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] k个一组翻转链表
 *
 * https://leetcode-cn.com/problems/reverse-nodes-in-k-group/description/
 *
 * algorithms
 * Hard (51.48%)
 * Likes:    201
 * Dislikes: 0
 * Total Accepted:    11.2K
 * Total Submissions: 21.6K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给出一个链表，每 k 个节点一组进行翻转，并返回翻转后的链表。
 *
 * k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么将最后剩余节点保持原有顺序。
 *
 * 示例 :
 *
 * 给定这个链表：1->2->3->4->5
 *
 * 当 k = 2 时，应当返回: 2->1->4->3->5
 *
 * 当 k = 3 时，应当返回: 3->2->1->4->5
 *
 * 说明 :
 *
 *
 * 你的算法只能使用常数的额外空间。
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
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
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	fast, slow := new(ListNode), new(ListNode)
	fast.Next = head
	slow.Next = head
	headers := make([]*ListNode, 0)
	for fast != nil {
		cnt := k
		for cnt > 0 && fast != nil {
			fast = fast.Next
			cnt--
		}
		if fast != nil {
			next := fast.Next
			header, tailer := reverse(slow.Next, fast)
			slow.Next = header
			tailer.Next = next
			headers = append(headers, header)
			slow = tailer
			fast = tailer
		}
	}
	if len(headers) > 0 {
		return headers[0]
	} else {
		return head
	}
}

func reverse(from, to *ListNode) (*ListNode, *ListNode) {
	if from == to {
		return from, to
	}
	head, tail := reverse(from.Next, to)
	from.Next = nil
	tail.Next = from
	return head, from
}

