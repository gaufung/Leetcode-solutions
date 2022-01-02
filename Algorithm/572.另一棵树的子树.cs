/*
 * @lc app=leetcode.cn id=572 lang=csharp
 *
 * [572] 另一棵树的子树
 *
 * https://leetcode-cn.com/problems/subtree-of-another-tree/description/
 *
 * algorithms
 * Easy (47.40%)
 * Likes:    617
 * Dislikes: 0
 * Total Accepted:    95.4K
 * Total Submissions: 201.2K
 * Testcase Example:  '[3,4,5,1,2]\n[4,1,2]'
 *
 * 
 * 
 * 给你两棵二叉树 root 和 subRoot 。检验 root 中是否包含和 subRoot 具有相同结构和节点值的子树。如果存在，返回 true
 * ；否则，返回 false 。
 * 
 * 二叉树 tree 的一棵子树包括 tree 的某个节点和这个节点的所有后代节点。tree 也可以看做它自身的一棵子树。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：root = [3,4,5,1,2], subRoot = [4,1,2]
 * 输出：true
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
 * 输出：false
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * root 树上的节点数量范围是 [1, 2000]
 * subRoot 树上的节点数量范围是 [1, 1000]
 * -10^4 
 * -10^4 
 * 
 * 
 * 
 * 
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     public int val;
 *     public TreeNode left;
 *     public TreeNode right;
 *     public TreeNode(int val=0, TreeNode left=null, TreeNode right=null) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
public class Solution {
    public bool IsSubtree(TreeNode root, TreeNode subRoot) {
        if (root == null && subRoot == null) {
            return true;
        }
        if (root == null && subRoot != null) {
            return false;
        }
        if (root != null && subRoot == null)
        {
            return false;
        }
        if (!IsEqual(root, subRoot)) {
            return  IsSubtree(root.left, subRoot) || IsSubtree(root.right, subRoot);
        }
        else{
            return IsSubtree(root.left, subRoot.left) && IsSubtree(root.right, subRoot.right);
        }
    }

    private bool IsEqual(TreeNode node1, TreeNode node2) {
        if (node1 == null && node2 == null) {
            return true;
        }
        if (node1 != null && node2 == null) {
            return false;
        }
        if (node1 == null && node2 != null) {
            return false;
        }

        return node1.val == node2.val && IsEqual(node1.left, node2.left) && IsEqual(node1.right, node2.right); 
        
    }
}
// @lc code=end

