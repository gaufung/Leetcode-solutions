import "fmt"

/*
 * @lc app=leetcode.cn id=459 lang=golang
 *
 * [459] 重复的子字符串
 *
 * https://leetcode-cn.com/problems/repeated-substring-pattern/description/
 *
 * algorithms
 * Easy (40.88%)
 * Likes:    95
 * Dislikes: 0
 * Total Accepted:    6.1K
 * Total Submissions: 14.9K
 * Testcase Example:  '"abab"'
 *
 * 给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。
 *
 * 示例 1:
 *
 *
 * 输入: "abab"
 *
 * 输出: True
 *
 * 解释: 可由子字符串 "ab" 重复两次构成。
 *
 *
 * 示例 2:
 *
 *
 * 输入: "aba"
 *
 * 输出: False
 *
 *
 * 示例 3:
 *
 *
 * 输入: "abcabcabcabc"
 *
 * 输出: True
 *
 * 解释: 可由子字符串 "abc" 重复四次构成。 (或者子字符串 "abcabc" 重复两次构成。)
 *
 *
 */
func repeatedSubstringPattern(s string) bool {
	n := len(s)
	if n <= 0 {
		return false
	}
	for i := 1; i < n; i++ {
		if repeat(s[i:], s[:i]) {
			return true
		}
	}
	return false
}

func repeat(text, pattern string) bool {
	fmt.Printf("%s, %s", text, pattern)
	n, m := len(text), len(pattern)
	if m > n {
		return false
	}
	if m == n {
		return text == pattern
	}
	if pattern != text[:m] {
		return false
	}
	return repeat(text[m:], pattern)
}

