/*
 * @lc app=leetcode.cn id=557 lang=csharp
 *
 * [557] 反转字符串中的单词 III
 *
 * https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/description/
 *
 * algorithms
 * Easy (74.37%)
 * Likes:    389
 * Dislikes: 0
 * Total Accepted:    190.3K
 * Total Submissions: 255.9K
 * Testcase Example:  `"Let's take LeetCode contest"`
 *
 * 给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
 * 
 * 
 * 
 * 示例：
 * 
 * 输入："Let's take LeetCode contest"
 * 输出："s'teL ekat edoCteeL tsetnoc"
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。
 * 
 * 
 */

// @lc code=start
public class Solution {
    public string ReverseWords(string s) {
        string[] words = s.Split(' ');
        for(int i = 0; i < words.Length; i++) {
            words[i] = ReverseWord(words[i]);
        }

        return string.Join(' ', words);
    }

    private string ReverseWord(string word){
        char[] digits = new char[word.Length];
        for(int i = 0; i < word.Length; i++){
            digits[i] = word[word.Length - i - 1];
        }

        return new string(digits);
    }
}
// @lc code=end

