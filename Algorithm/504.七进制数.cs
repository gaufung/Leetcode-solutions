/*
 * @lc app=leetcode.cn id=504 lang=csharp
 *
 * [504] 七进制数
 *
 * https://leetcode-cn.com/problems/base-7/description/
 *
 * algorithms
 * Easy (45.53%)
 * Likes:    20
 * Dislikes: 0
 * Total Accepted:    5.3K
 * Total Submissions: 11.4K
 * Testcase Example:  '100'
 *
 * 给定一个整数，将其转化为7进制，并以字符串形式输出。
 * 
 * 示例 1:
 * 
 * 
 * 输入: 100
 * 输出: "202"
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入: -7
 * 输出: "-10"
 * 
 * 
 * 注意: 输入范围是 [-1e7, 1e7] 。
 * 
 */
 using System.Text;
 using System.Linq;
public class Solution {
    public string ConvertToBase7(int num) {
        if(num == 0)
        {
            return "0";
        }
        if(num < 0) 
        {
            return "-" + ConvertToBase7(-1 * num);
        }
        StringBuilder sb = new StringBuilder();
        while(num > 0) 
        {
            sb.Append((num % 7).ToString());
            num /= 7;
        }
        return new string(sb.ToString().Reverse().ToArray());
    }
}

