/*
 * @lc app=leetcode.cn id=647 lang=golang
 *
 * [647] 回文子串
 *
 * https://leetcode-cn.com/problems/palindromic-substrings/description/
 *
 * algorithms
 * Medium (53.11%)
 * Likes:    90
 * Dislikes: 0
 * Total Accepted:    3.5K
 * Total Submissions: 6.6K
 * Testcase Example:  '"abc"'
 *
 * 给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
 *
 * 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被计为是不同的子串。
 *
 * 示例 1:
 *
 *
 * 输入: "abc"
 * 输出: 3
 * 解释: 三个回文子串: "a", "b", "c".
 *
 *
 * 示例 2:
 *
 *
 * 输入: "aaa"
 * 输出: 6
 * 说明: 6个回文子串: "a", "a", "a", "aa", "aa", "aaa".
 *
 *
 * 注意:
 *
 *
 * 输入的字符串长度不会超过1000。
 *
 *
 */
func countSubstrings(s string) int {
	n := len(s)
	if n <= 0 {
		return 0
	}
	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if valid(s[j : i+1]) {
				sum++
			}
		}
	}
	return sum

}

func valid(s string) bool {
	lo, hi := 0, len(s)-1
	for lo < hi {
		if s[lo] != s[hi] {
			return false
		}
		lo++
		hi--
	}
	return true
}

