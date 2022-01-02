/*
 * @lc app=leetcode.cn id=520 lang=csharp
 *
 * [520] 检测大写字母
 *
 * https://leetcode-cn.com/problems/detect-capital/description/
 *
 * algorithms
 * Easy (57.60%)
 * Likes:    196
 * Dislikes: 0
 * Total Accepted:    68.7K
 * Total Submissions: 119.3K
 * Testcase Example:  '"USA"'
 *
 * 我们定义，在以下情况时，单词的大写用法是正确的：
 * 
 * 
 * 全部字母都是大写，比如 "USA" 。
 * 单词中所有字母都不是大写，比如 "leetcode" 。
 * 如果单词不只含有一个字母，只有首字母大写， 比如 "Google" 。
 * 
 * 
 * 给你一个字符串 word 。如果大写用法正确，返回 true ；否则，返回 false 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：word = "USA"
 * 输出：true
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：word = "FlaG"
 * 输出：false
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= word.length <= 100
 * word 由小写和大写英文字母组成
 * 
 * 
 */

// @lc code=start
public class Solution {
    public bool DetectCapitalUse(string word) {
        if (word.Length == 1){
            return true;
        }

        if (word[0] >= 'a' && word[0] <= 'z')
        {
            //todo
            for(int i = 1; i < word.Length; i++)
            {
                if (word[i] >= 'a' && word[i] <= 'z')
                {
                    continue;
                }
                else
                {
                    return false;
                }
            }
            return true;
        }
        else 
        {
            if (word[1] >='a' && word[1] <= 'z')
            {
                for(int i = 2; i<word.Length; i++)
                {
                    if (word[i] >= 'a' && word[i] <= 'z')
                    {
                        continue;
                    }
                    else
                    {
                        return false;
                    }
                }
                return true;
            }
            else
            {
                for(int i = 2; i < word.Length; i++)
                {
                    if (word[i] >= 'A' && word[i] <='Z')
                    {
                        continue;
                    }
                    else
                    {
                        return false;
                    }
                }
                return true;
            }
        }
    }
}
// @lc code=end

