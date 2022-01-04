/*
 * @lc app=leetcode.cn id=640 lang=csharp
 *
 * [640] 求解方程
 *
 * https://leetcode-cn.com/problems/solve-the-equation/description/
 *
 * algorithms
 * Medium (42.66%)
 * Likes:    80
 * Dislikes: 0
 * Total Accepted:    11.4K
 * Total Submissions: 26.7K
 * Testcase Example:  '"x+5-3+x=6+x-2"'
 *
 * 求解一个给定的方程，将x以字符串"x=#value"的形式返回。该方程仅包含'+'，' - '操作，变量 x 和其对应系数。
 * 
 * 如果方程没有解，请返回“No solution”。
 * 
 * 如果方程有无限解，则返回“Infinite solutions”。
 * 
 * 如果方程中只有一个解，要保证返回值 x 是一个整数。
 * 
 * 示例 1：
 * 
 * 输入: "x+5-3+x=6+x-2"
 * 输出: "x=2"
 * 
 * 
 * 示例 2:
 * 
 * 输入: "x=x"
 * 输出: "Infinite solutions"
 * 
 * 
 * 示例 3:
 * 
 * 输入: "2x=x"
 * 输出: "x=0"
 * 
 * 
 * 示例 4:
 * 
 * 输入: "2x+3x-6x=x+2"
 * 输出: "x=-1"
 * 
 * 
 * 示例 5:
 * 
 * 输入: "x=x+2"
 * 输出: "No solution"
 * 
 * 
 */

// @lc code=start
public class Solution {
    public string SolveEquation(string equation) {
        string[] segements = equation.Split('=');
        var (leftA, leftB) = Parse(segements[0]);
        var (rightA, rightB) = Parse(segements[1]);
        var a = leftA - rightA;
        var b = rightB - leftB;

        if (a == 0 && b == 0 ){
            return "Infinite solutions";
        }
        if (a == 0 && b!=0) {
            return "No solution";
        }
        return $"x={b/a}";
    }
    
    private (int, int) Parse(string expression) {
        if (!expression.StartsWith('-') && !expression.StartsWith('+')) 
        {
            expression = "+" + expression;
        } 
        IList<string> tokens = SplitAndKeep(expression, new []{'+', '-'}).ToList();
        int a = 0; 
        int b = 0;
        int idx = 0;
        while(idx < tokens.Count) {
            if (tokens[idx] == "+" || tokens[idx] == "-") {
                if (tokens[idx+1].EndsWith("x")) {
                    if (tokens[idx+1] == "x") {
                        a += int.Parse(tokens[idx] + "1");
                    }
                    else {
                        a += int.Parse(tokens[idx] + tokens[idx+1].Substring(0, tokens[idx+1].Length-1));
                    }
                }else {
                    b += int.Parse(tokens[idx] + tokens[idx+1]);
                }
                idx+=2;
            }
        }
        return (a, b);
    }

    public static IEnumerable<string> SplitAndKeep(string s, char[] delims)
    {
        int start = 0, index;
        while ((index = s.IndexOfAny(delims, start)) != -1)
        {
            if(index-start > 0)
                yield return s.Substring(start, index - start);
            yield return s.Substring(index, 1);
            start = index + 1;
        }

        if (start < s.Length)
        {
            yield return s.Substring(start);
        }
    }
}
// @lc code=end

