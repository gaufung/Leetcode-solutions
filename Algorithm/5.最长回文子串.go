/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 *
 * https://leetcode-cn.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (25.19%)
 * Likes:    876
 * Dislikes: 0
 * Total Accepted:    59.7K
 * Total Submissions: 236.9K
 * Testcase Example:  '"babad"'
 *
 * 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
 *
 * 示例 1：
 *
 * 输入: "babad"
 * 输出: "bab"
 * 注意: "aba" 也是一个有效答案。
 *
 *
 * 示例 2：
 *
 * 输入: "cbbd"
 * 输出: "bb"
 *
 *
 */
func longestPalindrome(s string) string {
	length, result := 0, make([]rune, 0)
	buffer := []rune(s)
	for idx, _ := range s {
		current_length := palindromeLengthSymmetic(s, idx, idx+1)
		if current_length > length {
			result = buffer[idx-(current_length/2)+1 : idx+(current_length/2)+1]
			length = current_length
		}
		current_length = palindromeLengthNonSymmetric(s, idx)
		if current_length > length {
			result = buffer[idx-(current_length-1)/2 : idx+(current_length-1)/2+1]
			length = current_length
		}
	}
	return string(result)
}
func palindromeLengthSymmetic(s string, left, right int) int {
	result := 0
	for left > -1 && right < len(s) {
		if s[left] == s[right] {
			left--
			right++
			result += 2
		} else {
			break
		}
	}
	return result
}

func palindromeLengthNonSymmetric(s string, middle int) int {
	result := 1
	return result + palindromeLengthSymmetic(s, middle-1, middle+1)
}

