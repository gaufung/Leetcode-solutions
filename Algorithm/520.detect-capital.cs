/*
 * @lc app=leetcode id=520 lang=csharp
 *
 * [520] Detect Capital
 *
 * https://leetcode.com/problems/detect-capital/description/
 *
 * algorithms
 * Easy (54.26%)
 * Likes:    974
 * Dislikes: 313
 * Total Accepted:    208.8K
 * Total Submissions: 384.9K
 * Testcase Example:  '"USA"'
 *
 * We define the usage of capitals in a word to be right when one of the
 * following cases holds:
 * 
 * 
 * All letters in this word are capitals, like "USA".
 * All letters in this word are not capitals, like "leetcode".
 * Only the first letter in this word is capital, like "Google".
 * 
 * 
 * Given a string word, return true if the usage of capitals in it is right.
 * 
 * 
 * Example 1:
 * Input: word = "USA"
 * Output: true
 * Example 2:
 * Input: word = "FlaG"
 * Output: false
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= word.length <= 100
 * word consists of lowercase and uppercase English letters.
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

