/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
    slow, fast := head, head
    for n > 0 {
        fast = fast.Next
        n--
    }
    for {
        if  fast == nil || fast.Next == nil {
            break
        }else{
            fast = fast.Next
            slow = slow.Next
        }
    }
    if fast == nil {
        return slow.Next
    }else{
        slow.Next = slow.Next.Next
        return head
    }   
}