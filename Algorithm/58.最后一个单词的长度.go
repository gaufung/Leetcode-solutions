/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 *
 * https://leetcode-cn.com/problems/length-of-last-word/description/
 *
 * algorithms
 * Easy (29.99%)
 * Likes:    107
 * Dislikes: 0
 * Total Accepted:    28.4K
 * Total Submissions: 94.8K
 * Testcase Example:  '"Hello World"'
 *
 * 给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。
 *
 * 如果不存在最后一个单词，请返回 0 。
 *
 * 说明：一个单词是指由字母组成，但不包含任何空格的字符串。
 *
 * 示例:
 *
 * 输入: "Hello World"
 * 输出: 5
 *
 *
 */
func lengthOfLastWord(s string) int {
	n := len(s)
	if n <= 0 {
		return 0
	}
	index1 := notEmpty(s, n-1)
	if index1 == -1 {
		return 0
	}
	index2 := empty(s, index1)
	return index1 - index2
}

func notEmpty(s string, from int) int {
	for from >= 0 && s[from] == ' ' {
		from--
	}
	return from
}
func empty(s string, from int) int {
	for from >= 0 && s[from] != ' ' {
		from--
	}
	return from
}

