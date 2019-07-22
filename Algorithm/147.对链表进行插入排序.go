/*
 * @lc app=leetcode.cn id=147 lang=golang
 *
 * [147] 对链表进行插入排序
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}
	first, last := head, head
	for last.Next.Next!= nil {
		if last.Next.Val >= last.Val {
			last = last.Next
		}else{
			node := last.Next
			last.Next = node.Next
			if node.Val <= first.Val {
				node.Next = first
				first = node
			}else{
				cur := first
				for cur.Next.Val < node.Val {
					cur = cur.Next
				}
				node.Next = cur.Next
				cur.Next = node
			}
		}
	}
	if last.Next.Val >= last.Val {
		return first
	}
	node := last.Next
	last.Next = nil
	if node.Val <= first.Val {
		node.Next = first
		first = node
	}else{
		cur := first
		for cur.Next.Val < node.Val {
			cur = cur.Next
		}
		node.Next = cur.Next
		cur.Next = node
	}
	return first

}

