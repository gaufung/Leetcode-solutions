/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] 反转字符串中的元音字母
 *
 * https://leetcode-cn.com/problems/reverse-vowels-of-a-string/description/
 *
 * algorithms
 * Easy (46.79%)
 * Likes:    47
 * Dislikes: 0
 * Total Accepted:    10.2K
 * Total Submissions: 21.7K
 * Testcase Example:  '"hello"'
 *
 * 编写一个函数，以字符串作为输入，反转该字符串中的元音字母。
 *
 * 示例 1:
 *
 * 输入: "hello"
 * 输出: "holle"
 *
 *
 * 示例 2:
 *
 * 输入: "leetcode"
 * 输出: "leotcede"
 *
 * 说明:
 * 元音字母不包含字母"y"。
 *
 */
func reverseVowels(s string) string {
	bytes := []byte(s)
	i, j := 0, len(bytes)-1
	for i < j {
		for i < j && !isVowels(bytes[i]) {
			i++
		}
		for i < j && !isVowels(bytes[j]) {
			j--
		}
		if i < j {
			bytes[i], bytes[j] = bytes[j], bytes[i]
			i++
			j--
		} else {
			break
		}
	}
	return string(bytes)
}

func isVowels(ch byte) bool {
	//  a, e, i, o, u
	return ch == 'a' || ch == 'A' || ch == 'e' || ch == 'E' || ch == 'I' || ch == 'i' || ch == 'o' || ch == 'O' || ch == 'u' || ch == 'U'
}

