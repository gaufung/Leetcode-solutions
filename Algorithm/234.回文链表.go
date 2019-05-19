/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 *
 * https://leetcode-cn.com/problems/palindrome-linked-list/description/
 *
 * algorithms
 * Easy (36.73%)
 * Likes:    206
 * Dislikes: 0
 * Total Accepted:    23.5K
 * Total Submissions: 64K
 * Testcase Example:  '[1,2]'
 *
 * 请判断一个链表是否为回文链表。
 *
 * 示例 1:
 *
 * 输入: 1->2
 * 输出: false
 *
 * 示例 2:
 *
 * 输入: 1->2->2->1
 * 输出: true
 *
 *
 * 进阶：
 * 你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head.Next
	cnt := 2
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
		cnt++
		if fast.Next != nil {
			fast = fast.Next
			cnt++
		}
	}
	head2 := slow.Next
	slow.Next = nil
	head1, _ := reverse(head)
	if cnt%2 != 0 {
		head1 = head1.Next
	}
	for head1 != nil && head2 != nil {
		if head1.Val != head2.Val {
			return false
		}
		head1 = head1.Next
		head2 = head2.Next
	}
	return true

}

func reverse(head *ListNode) (*ListNode, *ListNode) {
	if head.Next == nil {
		return head, head
	}
	h, t := reverse(head.Next)
	t.Next = head
	head.Next = nil
	return h, head
}

