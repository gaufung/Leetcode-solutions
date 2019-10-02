/*
 * @lc app=leetcode.cn id=430 lang=csharp
 *
 * [430] 扁平化多级双向链表
 *
 * https://leetcode-cn.com/problems/flatten-a-multilevel-doubly-linked-list/description/
 *
 * algorithms
 * Medium (41.18%)
 * Likes:    43
 * Dislikes: 0
 * Total Accepted:    3.6K
 * Total Submissions: 8.5K
 * Testcase Example:  '{"$id":"1","child":null,"next":{"$id":"2","child":null,"next":{"$id":"3","child":{"$id":"7","child":null,"next":{"$id":"8","child":{"$id":"11","child":null,"next":{"$id":"12","child":null,"next":null,"prev":{"$ref":"11"},"val":12},"prev":null,"val":11},"next":{"$id":"9","child":null,"next":{"$id":"10","child":null,"next":null,"prev":{"$ref":"9"},"val":10},"prev":{"$ref":"8"},"val":9},"prev":{"$ref":"7"},"val":8},"prev":null,"val":7},"next":{"$id":"4","child":null,"next":{"$id":"5","child":null,"next":{"$id":"6","child":null,"next":null,"prev":{"$ref":"5"},"val":6},"prev":{"$ref":"4"},"val":5},"prev":{"$ref":"3"},"val":4},"prev":{"$ref":"2"},"val":3},"prev":{"$ref":"1"},"val":2},"prev":null,"val":1}'
 *
 * 
 * 您将获得一个双向链表，除了下一个和前一个指针之外，它还有一个子指针，可能指向单独的双向链表。这些子列表可能有一个或多个自己的子项，依此类推，生成多级数据结构，如下面的示例所示。
 * 
 * 扁平化列表，使所有结点出现在单级双链表中。您将获得列表第一级的头部。
 * 
 * 
 * 
 * 示例:
 * 
 * 输入:
 * ⁠1---2---3---4---5---6--NULL
 * ⁠        |
 * ⁠        7---8---9---10--NULL
 * ⁠            |
 * ⁠            11--12--NULL
 * 
 * 输出:
 * 1-2-3-7-8-11-12-9-10-4-5-6-NULL
 * 
 * 
 * 
 * 
 * 以上示例的说明:
 * 
 * 给出以下多级双向链表:
 * 
 * 
 * 
 * 
 * 
 * 我们应该返回如下所示的扁平双向链表:
 * 
 * 
 * 
 */

// @lc code=start
/*
// Definition for a Node.
public class Node {
    public int val;
    public Node prev;
    public Node next;
    public Node child;

    public Node(){}
    public Node(int _val,Node _prev,Node _next,Node _child) {
        val = _val;
        prev = _prev;
        next = _next;
        child = _child;
}
*/
using System.Collections.Generic;
public class Solution {
    public Node Flatten(Node head) {
        if(head == null)
        {
            return null;
        }

        Stack<Node> next = new Stack<Node>();
        next.Push(head);
        Node first = new Node(0, null, null, null);
        Node cur = first;
        while(next.Count > 0)
        {
            Node node= next.Pop();
            cur.next = new Node(node.val, cur, null, null);
            if(node.child == null && node.next != null)
            {
                next.Push(node.next);
            }
            else if(node.child!=null && node.next == null)
            {
                next.Push(node.child);
            }
            else if(node.child!=null && node.next !=null)
            {
                next.Push(node.next);
                next.Push(node.child);
            }
            cur = cur.next;
        }
        return first.next;
    }
}
// @lc code=end

