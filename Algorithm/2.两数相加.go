/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	first, second := l1, l2
	carry := 0
	result := &ListNode{(first.Val + second.Val + carry) % 10, nil}
	carry = (first.Val + second.Val + carry) / 10
	cur := result
	for first.Next != nil && second.Next != nil {
		first = first.Next
		second = second.Next
		cur.Next = &ListNode{(first.Val + second.Val + carry) % 10, nil}
		carry = (first.Val + second.Val + carry) / 10
		cur = cur.Next
	}
	cur, carry = addOneChain(first, cur, carry)
	cur, carry = addOneChain(second, cur, carry)
	if carry == 1 {
		cur.Next = &ListNode{1, nil}
	}
	return result

}

func addOneChain(li, cur *ListNode, carry int) (*ListNode, int) {
	for li.Next != nil {
		li = li.Next
		cur.Next = &ListNode{(li.Val + carry) % 10, nil}
		carry = (li.Val + carry) / 10
		cur = cur.Next
	}
	return cur, carry
}

