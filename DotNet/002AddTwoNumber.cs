/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     public int val;
 *     public ListNode next;
 *     public ListNode(int x) { val = x; }
 * }
 */
public class Solution {
    public ListNode AddTwoNumbers(ListNode l1, ListNode l2) {
        ListNode first = l1;
        ListNode second = l2;
        int carry = 0;
        ListNode result = new ListNode((first.val+second.val+carry)%10);
        carry = (first.val+second.val+carry) / 10;
        ListNode cur = result;
        while(first.next != null && second.next != null){
            first = first.next;
            second = second.next;
            cur.next = new ListNode((first.val+second.val+carry)%10);
            carry = (first.val+second.val+carry) / 10;
            cur = cur.next;
        }
        while(first.next!=null){
            first = first.next;
            cur.next = new ListNode((first.val+carry)%10);
            carry = (first.val+carry) / 10;
            cur = cur.next;
        }
        while(second.next!=null){
            second = second.next;
            cur.next = new ListNode((second.val+carry)%10);
            carry = (second.val + carry) / 10;
            cur = cur.next;
        }
        if(carry==1){
            cur.next = new ListNode(1);
        }
        return result;
        
    }
}