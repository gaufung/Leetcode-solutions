/*
 * @lc app=leetcode.cn id=385 lang=csharp
 *
 * [385] 迷你语法分析器
 *
 * https://leetcode-cn.com/problems/mini-parser/description/
 *
 * algorithms
 * Medium (38.82%)
 * Likes:    10
 * Dislikes: 0
 * Total Accepted:    812
 * Total Submissions: 2.1K
 * Testcase Example:  '"324"'
 *
 * 给定一个用字符串表示的整数的嵌套列表，实现一个解析它的语法分析器。
 * 
 * 列表中的每个元素只可能是整数或整数嵌套列表
 * 
 * 提示：你可以假定这些字符串都是格式良好的：
 * 
 * 
 * 字符串非空
 * 字符串不包含空格
 * 字符串只包含数字0-9, [, - ,, ]
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 给定 s = "324",
 * 
 * 你应该返回一个 NestedInteger 对象，其中只包含整数值 324。
 * 
 * 
 * 
 * 
 * 示例 2：
 * 
 * 
 * 给定 s = "[123,[456,[789]]]",
 * 
 * 返回一个 NestedInteger 对象包含一个有两个元素的嵌套列表：
 * 
 * 1. 一个 integer 包含值 123
 * 2. 一个包含两个元素的嵌套列表：
 * ⁠   i.  一个 integer 包含值 456
 * ⁠   ii. 一个包含一个元素的嵌套列表
 * ⁠        a. 一个 integer 包含值 789
 * 
 * 
 * 
 * 
 */
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * interface NestedInteger {
 *
 *     // Constructor initializes an empty nested list.
 *     public NestedInteger();
 *
 *     // Constructor initializes a single integer.
 *     public NestedInteger(int value);
 *
 *     // @return true if this NestedInteger holds a single integer, rather than a nested list.
 *     bool IsInteger();
 *
 *     // @return the single integer that this NestedInteger holds, if it holds a single integer
 *     // Return null if this NestedInteger holds a nested list
 *     int GetInteger();
 *
 *     // Set this NestedInteger to hold a single integer.
 *     public void SetInteger(int value);
 *
 *     // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 *     public void Add(NestedInteger ni);
 *
 *     // @return the nested list that this NestedInteger holds, if it holds a nested list
 *     // Return null if this NestedInteger holds a single integer
 *     IList<NestedInteger> GetList();
 * }
 */
 using System.Collections.Generic;
public class Solution {
    public NestedInteger Deserialize(string s) {
        if(!s.StartsWith('['))
        {
            int i = 0;
            int number = ReadNumber(s, ref i);
            return new NestedInteger(number);
        }
        s = "[" + s + "]";
        Stack<char> ops = new Stack<char>();
        Stack<int> operands = new Stack<int>();
        int pos = 0;
        NestedInteger result = null;
        while(pos < s.Length)
        {
            if(s[pos]=='[') 
            {
                ops.Push('[');
                pos++;
            }
            else if (IsDigit(s[pos]) || s[pos]=='-')
            {
                int number = ReadNumber(s, ref pos);
                operands.Push(number);
            }
            else if (s[pos]==']')
            {
                if(operands.Count == 0) {
                    var temp = new NestedInteger();
                    if(result == null)
                    {
                        result = temp;
                    }
                    else
                    {
                        temp.Add(result);
                        result = temp;
                    }
                    pos++;
                    continue;
                }
                int number = operands.Pop();
                ops.Pop();
                if(result == null) 
                {
                    result = new NestedInteger(number);
                }
                else
                {
                    var temp = new NestedInteger();
                    temp.SetInteger(number);
                    temp.Add(result);
                    result = temp;
                }
                pos++;
            }
            else
            {
                pos++;
            }
        }
        return result;
    }
    private bool IsDigit(char ch)
    {
        return ch >= '0' && ch <= '9';
    }

    private int ReadNumber(string s, ref int pos)
    {
        bool isNegative = s[pos] == '-';
        int result = 0;
        if(isNegative)
        {
            pos++;
        }
        while(pos < s.Length && IsDigit(s[pos]))
        {
            result = result * 10 + (int)(s[pos]-'0');
            pos++;
        }
        return isNegative? -1 * result : result;
    }
}

