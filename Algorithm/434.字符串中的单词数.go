/*
 * @lc app=leetcode.cn id=434 lang=golang
 *
 * [434] 字符串中的单词数
 *
 * https://leetcode-cn.com/problems/number-of-segments-in-a-string/description/
 *
 * algorithms
 * Easy (30.67%)
 * Likes:    34
 * Dislikes: 0
 * Total Accepted:    6K
 * Total Submissions: 19.3K
 * Testcase Example:  '"Hello, my name is John"'
 *
 * 统计字符串中的单词个数，这里的单词指的是连续的不是空格的字符。
 *
 * 请注意，你可以假定字符串里不包括任何不可打印的字符。
 *
 * 示例:
 *
 * 输入: "Hello, my name is John"
 * 输出: 5
 *
 *
 */
func countSegments(s string) int {
	cnt, n := 0, len(s)
	i := 0
	for i < n {
		if s[i] == ' ' {
			i++
		} else {
			cnt++
			for i < n && s[i] != ' ' {
				i++
			}
		}
	}
	return cnt
}

