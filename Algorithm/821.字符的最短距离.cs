/*
 * @lc app=leetcode.cn id=821 lang=csharp
 *
 * [821] 字符的最短距离
 *
 * https://leetcode-cn.com/problems/shortest-distance-to-a-character/description/
 *
 * algorithms
 * Easy (64.21%)
 * Likes:    84
 * Dislikes: 0
 * Total Accepted:    7.3K
 * Total Submissions: 11.3K
 * Testcase Example:  '"loveleetcode"\n"e"'
 *
 * 给定一个字符串 S 和一个字符 C。返回一个代表字符串 S 中每个字符到字符串 S 中的字符 C 的最短距离的数组。
 * 
 * 示例 1:
 * 
 * 
 * 输入: S = "loveleetcode", C = 'e'
 * 输出: [3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0]
 * 
 * 
 * 说明:
 * 
 * 
 * 字符串 S 的长度范围为 [1, 10000]。
 * C 是一个单字符，且保证是字符串 S 里的字符。
 * S 和 C 中的所有字母均为小写字母。
 * 
 * 
 */

using System;
public class Solution {
    public int[] ShortestToChar(string S, char C) {
        int[] result = new int[S.Length];
        for(int i=0; i < result.Length; i++)
        {
            result[i] = int.MaxValue;
        }
        for(int i=0; i < S.Length; i++)
        {
            if(S[i] == C)
            {
                Update(result, i);
            }
        }
        return result;
    }

    private void Update(int[] distance, int index)
    {
        for(int j =index; j >=0; j--)
        {
            distance[j] = Math.Min(distance[j], index-j);
        }
        for(int j=index; j<distance.Length; j++)
        {
            distance[j] = Math.Min(distance[j], j-index);
        }
    }
}

