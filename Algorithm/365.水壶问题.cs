/*
 * @lc app=leetcode.cn id=365 lang=csharp
 *
 * [365] 水壶问题
 *
 * https://leetcode-cn.com/problems/water-and-jug-problem/description/
 *
 * algorithms
 * Medium (27.96%)
 * Likes:    27
 * Dislikes: 0
 * Total Accepted:    2.1K
 * Total Submissions: 7.3K
 * Testcase Example:  '3\n5\n4'
 *
 * 有两个容量分别为 x升 和 y升 的水壶以及无限多的水。请判断能否通过使用这两个水壶，从而可以得到恰好 z升 的水？
 * 
 * 如果可以，最后请用以上水壶中的一或两个来盛放取得的 z升 水。
 * 
 * 你允许：
 * 
 * 
 * 装满任意一个水壶
 * 清空任意一个水壶
 * 从一个水壶向另外一个水壶倒水，直到装满或者倒空
 * 
 * 
 * 示例 1: (From the famous "Die Hard" example)
 * 
 * 输入: x = 3, y = 5, z = 4
 * 输出: True
 * 
 * 
 * 示例 2:
 * 
 * 输入: x = 2, y = 6, z = 5
 * 输出: False
 * 
 * 
 */
 using System.Collections.Generic;
 using System;
public class Solution {

    private ISet<Status> nodes = new HashSet<Status>();
    public bool CanMeasureWater(int x, int y, int z) {
        nodes = new HashSet<Status>();
        var statuses = new List<Status>();
        statuses.Add(new Status(0, 0));
        while(statuses.Count > 0)
        {
            var nextNodes = new List<Status>();
            foreach (var status in statuses)
            {
                var currentX = status.X;
                var currentY = status.Y;
                if(currentX==z || currentY == z || currentX+currentY == z)
                {
                    return true;
                }
                var node1 = new Status(x, currentY);
                var node2 = new Status(currentX, y);
                var total = currentX + currentY;
                var node3 = new Status(total >= x? x : total, total >= x? total - x : 0);
                var node4 = new Status(total >= y? total - y: 0, total >= y? y: total);

                var node5 = new Status(0, currentY);
                var node6 = new Status(currentX, 0);
                var newNodes = new Status[] {node1, node2, node3, node4, node5, node6};
                foreach(var node in newNodes)
                {
                    if(!nodes.Contains(node))
                    {
                        nodes.Add(node);
                        nextNodes.Add(node);
                    }
                }
            }
            statuses = nextNodes;
        }
        return false;
    }
}

public class Status 
{
    public int X { get; set; }

    public int Y { get; set; }

    public Status(int x, int y)
    {
        X = x;
        Y = y;
    }

    public override bool Equals(object obj)
    {
        Status other = obj as Status;
        if(other == null) 
        {
            return false;
        }
        return other.X == this.X && other.Y == this.Y;
    }

    public override int GetHashCode()
    {
        return X.GetHashCode() ^ Y.GetHashCode();
    }
}

