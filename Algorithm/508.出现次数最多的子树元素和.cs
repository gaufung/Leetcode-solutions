using System.Collections.Generic;
using System.Linq;
/*
 * @lc app=leetcode.cn id=508 lang=csharp
 *
 * [508] 出现次数最多的子树元素和
 *
 * https://leetcode-cn.com/problems/most-frequent-subtree-sum/description/
 *
 * algorithms
 * Medium (58.19%)
 * Likes:    33
 * Dislikes: 0
 * Total Accepted:    1.8K
 * Total Submissions: 3.2K
 * Testcase Example:  '[5,2,-3]'
 *
 * 
 * 给出二叉树的根，找出出现次数最多的子树元素和。一个结点的子树元素和定义为以该结点为根的二叉树上所有结点的元素之和（包括结点本身）。然后求出出现次数最多的子树元素和。如果有多个元素出现的次数相同，返回所有出现次数最多的元素（不限顺序）。
 * 
 * 
 * 
 * 示例 1
 * 输入:
 * 
 * ⁠ 5
 * ⁠/  \
 * 2   -3
 * 
 * 
 * 返回 [2, -3, 4]，所有的值均只出现一次，以任意顺序返回所有值。
 * 
 * 示例 2
 * 输入:
 * 
 * ⁠ 5
 * ⁠/  \
 * 2   -5
 * 
 * 
 * 返回 [2]，只有 2 出现两次，-5 只出现 1 次。
 * 
 * 
 * 
 * 提示： 假设任意子树元素和均可以用 32 位有符号整数表示。
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
    private IDictionary<int, int> counter;
    public int[] FindFrequentTreeSum(TreeNode root) {
        if(root==null)
        {
            return new int[]{};
        }
        counter = new Dictionary<int, int>();
        travel(root);
        var max = counter.Select(KeyValuePair=>KeyValuePair.Value).Max();
        return counter.Where(pair=>pair.Value==max).Select(pair=>pair.Key).ToArray();
    }

    private int travel(TreeNode node) 
    {
        if(node.left==null && node.right==null)
        {
            if(this.counter.ContainsKey(node.val))
            {
                this.counter[node.val] += 1;
            }
            else
            {
                this.counter.Add(node.val, 1);
            }
            return node.val;
        }
        if(node.left==null && node.right!=null)
        {
            var sum = travel(node.right) + node.val;
            if(this.counter.ContainsKey(sum))
            {
                this.counter[sum] +=1;
            }
            else
            {
                this.counter.Add(sum, 1);
            }
            return sum;
        }
        if(node.left!=null && node.right==null)
        {
            var sum = travel(node.left) + node.val;
            if(this.counter.ContainsKey(sum))
            {
                this.counter[sum] += 1;
            }
            else
            {
                this.counter.Add(sum, 1);
            }
            return sum;
        }
        else
        {
             var sum = node.val + travel(node.left) + travel(node.right);
        if(this.counter.ContainsKey(sum))
        {
            this.counter[sum] +=1;
        }
        else
        {
            this.counter.Add(sum, 1);
        }
        return sum;
        }
       
    }
}

