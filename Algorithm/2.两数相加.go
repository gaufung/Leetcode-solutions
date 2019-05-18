/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 *
 * https://leetcode-cn.com/problems/add-two-numbers/description/
 *
 * algorithms
 * Medium (34.08%)
 * Likes:    2402
 * Dislikes: 0
 * Total Accepted:    128.5K
 * Total Submissions: 377.1K
 * Testcase Example:  '[2,4,3]\n[5,6,4]'
 *
 * 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
 *
 * 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
 *
 * 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
 *
 * 示例：
 *
 * 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
 * 输出：7 -> 0 -> 8
 * 原因：342 + 465 = 807
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

