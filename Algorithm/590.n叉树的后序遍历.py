#
# @lc app=leetcode.cn id=590 lang=python
#
# [590] N叉树的后序遍历
#
# https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/description/
#
# algorithms
# Easy (68.96%)
# Likes:    29
# Dislikes: 0
# Total Accepted:    8.1K
# Total Submissions: 11.7K
# Testcase Example:  '{"$id":"1","children":[{"$id":"2","children":[{"$id":"5","children":[],"val":5},{"$id":"6","children":[],"val":6}],"val":3},{"$id":"3","children":[],"val":2},{"$id":"4","children":[],"val":4}],"val":1}'
#
# 给定一个 N 叉树，返回其节点值的后序遍历。
# 
# 例如，给定一个 3叉树 :
# 
# 
# 
# 
# 
# 
# 
# 返回其后序遍历: [5,6,3,2,4,1].
# 
# 
# 
# 说明: 递归法很简单，你可以使用迭代法完成此题吗?
#
"""
# Definition for a Node.
class Node(object):
    def __init__(self, val, children):
        self.val = val
        self.children = children
"""
class Solution(object):
    def postorder(self, root):
        """
        :type root: Node
        :rtype: List[int]
        """
        result = []
        if root is None:
            return result
        if root.children is not None:
            for node in root.children:
                next = self.postorder(node)
                if next:
                    result.extend(next)
        result.append(root.val)
        return result
        

