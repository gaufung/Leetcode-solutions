/*
 * @lc app=leetcode.cn id=445 lang=golang
 *
 * [445] 两数相加 II
 *
 * https://leetcode-cn.com/problems/add-two-numbers-ii/description/
 *
 * algorithms
 * Medium (48.66%)
 * Likes:    63
 * Dislikes: 0
 * Total Accepted:    4.1K
 * Total Submissions: 8.5K
 * Testcase Example:  '[7,2,4,3]\n[5,6,4]'
 *
 * 给定两个非空链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储单个数字。将这两数相加会返回一个新的链表。
 *
 *
 *
 * 你可以假设除了数字 0 之外，这两个数字都不会以零开头。
 *
 * 进阶:
 *
 * 如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。
 *
 * 示例:
 *
 *
 * 输入: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
 * 输出: 7 -> 8 -> 0 -> 7
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	h1, _ := reverse(l1)
	h2, _ := reverse(l2)
	h := add(h1, h2)
	res, _ := reverse(h)
	return res
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

func add(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	sum := l1.Val + l2.Val
	val := sum % 10
	carrier := sum / 10
	head := &ListNode{
		Val: val,
	}
	l1 = l1.Next
	l2 = l2.Next
	cur := head
	for l1 != nil && l2 != nil {
		sum = l1.Val + l2.Val + carrier
		val = sum % 10
		carrier = sum / 10
		cur.Next = &ListNode{
			Val: val,
		}
		cur = cur.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		sum = l1.Val + carrier
		val = sum % 10
		carrier = sum / 10
		cur.Next = &ListNode{
			Val: val,
		}
		cur = cur.Next
		l1 = l1.Next
	}
	for l2 != nil {
		sum = l2.Val + carrier
		val = sum % 10
		carrier = sum / 10
		cur.Next = &ListNode{
			Val: val,
		}
		cur = cur.Next
		l2 = l2.Next
	}
	if carrier != 0 {
		cur.Next = &ListNode{
			Val: carrier,
		}
	}
	return head
}



