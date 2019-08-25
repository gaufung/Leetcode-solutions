/*
 * @lc app=leetcode.cn id=623 lang=csharp
 *
 * [623] 在二叉树中增加一行
 *
 * https://leetcode-cn.com/problems/add-one-row-to-tree/description/
 *
 * algorithms
 * Medium (47.82%)
 * Likes:    36
 * Dislikes: 0
 * Total Accepted:    1.6K
 * Total Submissions: 3.3K
 * Testcase Example:  '[4,2,6,3,1,5]\n1\n2'
 *
 * 给定一个二叉树，根节点为第1层，深度为 1。在其第 d 层追加一行值为 v 的节点。
 * 
 * 添加规则：给定一个深度值 d （正整数），针对深度为 d-1 层的每一非空节点 N，为 N 创建两个值为 v 的左子树和右子树。
 * 
 * 将 N 原先的左子树，连接为新节点 v 的左子树；将 N 原先的右子树，连接为新节点 v 的右子树。
 * 
 * 如果 d 的值为 1，深度 d - 1 不存在，则创建一个新的根节点 v，原先的整棵树将作为 v 的左子树。
 * 
 * 示例 1:
 * 
 * 
 * 输入: 
 * 二叉树如下所示:
 * ⁠      4
 * ⁠    /   \
 * ⁠   2     6
 * ⁠  / \   / 
 * ⁠ 3   1 5   
 * 
 * v = 1
 * 
 * d = 2
 * 
 * 输出: 
 * ⁠      4
 * ⁠     / \
 * ⁠    1   1
 * ⁠   /     \
 * ⁠  2       6
 * ⁠ / \     / 
 * ⁠3   1   5   
 * 
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入: 
 * 二叉树如下所示:
 * ⁠     4
 * ⁠    /   
 * ⁠   2    
 * ⁠  / \   
 * ⁠ 3   1    
 * 
 * v = 1
 * 
 * d = 3
 * 
 * 输出: 
 * ⁠     4
 * ⁠    /   
 * ⁠   2
 * ⁠  / \    
 * ⁠ 1   1
 * ⁠/     \  
 * 3       1
 * 
 * 
 * 注意:
 * 
 * 
 * 输入的深度值 d 的范围是：[1，二叉树最大深度 + 1]。
 * 输入的二叉树至少有一个节点。
 * 
 * 
 */
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     public int val;
 *     public TreeNode left;
 *     public TreeNode right;
 *     public TreeNode(int x) { val = x; }
 * }
 */
public class Solution {
    public TreeNode AddOneRow(TreeNode root, int v, int d) {
        if(d == 1)
        {
            TreeNode node=  new TreeNode(v);
            node.left = root;
            return node;
        }
        else
        {
            AddOneRow(root, v, d, 1);
            return root;
        }
    }

    private void AddOneRow(TreeNode node, int v, int d, int currentLevel)
    {
        if(currentLevel == d - 1)
        {
            TreeNode leftNode = new TreeNode(v);
            leftNode.left = node.left;
            node.left = leftNode;
            TreeNode rightNode = new TreeNode(v);
            rightNode.right = node.right;
            node.right = rightNode;
            return;
        }
        if(node.left != null)
        {
            AddOneRow(node.left, v, d, currentLevel+1);
        }
        if(node.right != null)
        {
            AddOneRow(node.right, v, d, currentLevel+1);
        }
    }
}

