# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        first = l1
        second = l2
        carry = 0
        result = ListNode((first.val+second.val+carry)%10)
        carry = (first.val+second.val+carry) // 10
        cur = result
        while first.next and second.next:
            first, second = first.next, second.next
            cur.next = ListNode((first.val+second.val+carry)%10)
            carry = (first.val+second.val+carry) // 10
            cur = cur.next
            
        cur,carry = self.addOneChain(first, cur, carry)
        cur,carry = self.addOneChain(second,cur, carry)
        if carry:
            cur.next = ListNode(1)
        return result
    def addOneChain(self, link, cur, carry):
        while link.next:
            link = link.next
            cur.next = ListNode((link.val + carry)%10)
            carry = (link.val + carry) // 10
            cur=cur.next
        return cur, carry