/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 *
 * https://leetcode-cn.com/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (39.55%)
 * Likes:    83
 * Dislikes: 0
 * Total Accepted:    35.8K
 * Total Submissions: 90.5K
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
 *
 * 说明：本题中，我们将空字符串定义为有效的回文串。
 *
 * 示例 1:
 *
 * 输入: "A man, a plan, a canal: Panama"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: "race a car"
 * 输出: false
 *
 *
 */
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	for i, j := 0, len(s)-1; i < j; {
		if isSkip(s[i]) {
			i++
			continue
		}
		if isSkip(s[j]) {
			j--
			continue
		}
		if isEqual(s[i], s[j]) {
			i++
			j--
			continue
		} else {
			return false
		}
	}
	return true
}

func isSkip(b byte) bool {
	if b >= 'a' && b <= 'z' {
		return false
	}
	if b >= 'A' && b <= 'Z' {
		return false
	}
	if b >= '0' && b <= '9' {
		return false
	}
	return true
}

func isEqual(a, b byte) bool {
	if a == b {
		return true
	}
	if isLowLetter(a) && isUpperLetter(b) {
		return a-('a'-'A') == b
	}
	if isLowLetter(b) && isUpperLetter(a) {
		return b-('a'-'A') == a
	}
	return false
}

func isLowLetter(b byte) bool {
	return b >= 'a' && b <= 'z'
}
func isUpperLetter(b byte) bool {
	return b >= 'A' && b <= 'Z'
}
func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

