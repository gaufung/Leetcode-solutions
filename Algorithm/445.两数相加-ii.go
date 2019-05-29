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
	val1 := convert(l1)
	val2 := convert(l2)
	return convertBack(val1 + val2)
}

func convert(node *ListNode) int {
	result := 0
	for node != nil {
		result = result*10 + node.Val
		node = node.Next
	}
	return result
}

func convertBack(val int) *ListNode {
	first := new(ListNode)
	if val == 0 {
		node := &ListNode{Val: 0}
		first.Next = node
	} else {
		for val > 0 {
			digit := val % 10
			node := &ListNode{Val: digit}
			node.Next = first.Next
			first.Next = node
			val /= 10
		}
	}

	return first.Next
}

