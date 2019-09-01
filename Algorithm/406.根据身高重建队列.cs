/*
 * @lc app=leetcode.cn id=406 lang=csharp
 *
 * [406] 根据身高重建队列
 *
 * https://leetcode-cn.com/problems/queue-reconstruction-by-height/description/
 *
 * algorithms
 * Medium (61.78%)
 * Likes:    94
 * Dislikes: 0
 * Total Accepted:    2.8K
 * Total Submissions: 4.6K
 * Testcase Example:  '[[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]'
 *
 * 假设有打乱顺序的一群人站成一个队列。 每个人由一个整数对(h, k)表示，其中h是这个人的身高，k是排在这个人前面且身高大于或等于h的人数。
 * 编写一个算法来重建这个队列。
 * 
 * 注意：
 * 总人数少于1100人。
 * 
 * 示例
 * 
 * 
 * 输入:
 * [[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]
 * 
 * 输出:
 * [[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]
 * 
 * 
 */
using System;
using System.Collections.Generic;
public class Solution {
    public int[][] ReconstructQueue(int[][] people) {
       Array.Sort(people, new PeopleCompare());
       List<int[]> result = new List<int[]>();
       foreach (var p in people)
       {
           result.Insert(p[1], p);
       }
       return result.ToArray();
    }
}

internal class PeopleCompare : IComparer
{
    public int Compare(object x, object y)
    {
        int[] people1 = x as int[];
        int [] people2 = y as int[];
        if(people1[0] > people2[0])
        {
            return -1;
        }
        else if (people1[0] < people2[0])
        {
            return 1;
        }
        return people1[1].CompareTo(people2[1]);
    }
}

