/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 *
 * https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/description/
 *
 * algorithms
 * Medium (41.14%)
 * Likes:    106
 * Dislikes: 0
 * Total Accepted:    10.4K
 * Total Submissions: 25.3K
 * Testcase Example:  '[1,2,3,3,4,4,5]'
 *
 * 给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
 *
 * 示例 1:
 *
 * 输入: 1->2->3->3->4->4->5
 * 输出: 1->2->5
 *
 *
 * 示例 2:
 *
 * 输入: 1->1->1->2->3
 * 输出: 2->3
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	cur := head
	same := false
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			same = true
			cur = cur.Next
		} else {
			break
		}
	}
	if same {
		return deleteDuplicates(cur.Next)
	} else {
		head.Next = deleteDuplicates(head.Next)
		return head
	}
}

