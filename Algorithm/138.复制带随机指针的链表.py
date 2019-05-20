#
# @lc app=leetcode.cn id=138 lang=python
#
# [138] 复制带随机指针的链表
#
# https://leetcode-cn.com/problems/copy-list-with-random-pointer/description/
#
# algorithms
# Medium (30.28%)
# Likes:    97
# Dislikes: 0
# Total Accepted:    6.5K
# Total Submissions: 21.6K
# Testcase Example:  '{"$id":"1","next":{"$id":"2","next":null,"random":{"$ref":"2"},"val":2},"random":{"$ref":"2"},"val":1}'
#
# 给定一个链表，每个节点包含一个额外增加的随机指针，该指针可以指向链表中的任何节点或空节点。
# 
# 要求返回这个链表的深拷贝。 
# 
# 
# 
# 示例：
# 
# 
# 
# 输入：
# 
# {"$id":"1","next":{"$id":"2","next":null,"random":{"$ref":"2"},"val":2},"random":{"$ref":"2"},"val":1}
# 
# 解释：
# 节点 1 的值是 1，它的下一个指针和随机指针都指向节点 2 。
# 节点 2 的值是 2，它的下一个指针指向 null，随机指针指向它自己。
# 
# 
# 
# 
# 提示：
# 
# 
# 你必须返回给定头的拷贝作为对克隆列表的引用。
# 
# 
#
"""
# Definition for a Node.
class Node(object):
    def __init__(self, val, next, random):
        self.val = val
        self.next = next
        self.random = random
"""
class Solution(object):
    def copyRandomList(self, head):
        """
        :type head: Node
        :rtype: Node
        """
        if head is None:
            return None
        m = dict()
        cur = head
        while cur is not None:
            m[cur] = Node(cur.val, None, None)
            cur = cur.next
        
        cur = head
        while cur is not None:
            if cur.next is not None:
                m[cur].next = m[cur.next]
            if cur.random is not None:
                m[cur].random = m[cur.random]
            cur = cur.next
        return m[head]
        

