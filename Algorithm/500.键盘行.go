/*
 * @lc app=leetcode.cn id=500 lang=golang
 *
 * [500] 键盘行
 *
 * https://leetcode-cn.com/problems/keyboard-row/description/
 *
 * algorithms
 * Easy (66.61%)
 * Likes:    46
 * Dislikes: 0
 * Total Accepted:    8K
 * Total Submissions: 12.1K
 * Testcase Example:  '["Hello","Alaska","Dad","Peace"]'
 *
 * 给定一个单词列表，只返回可以使用在键盘同一行的字母打印出来的单词。键盘如下图所示。
 *
 *
 *
 *
 *
 *
 *
 * 示例：
 *
 * 输入: ["Hello", "Alaska", "Dad", "Peace"]
 * 输出: ["Alaska", "Dad"]
 *
 *
 *
 *
 * 注意：
 *
 *
 * 你可以重复使用键盘上同一字符。
 * 你可以假设输入的字符串将只包含字母。
 *
 */
func findWords(words []string) []string {
	line1 := "qwertyuiopQWERTYUIOP"
	line2 := "asdfghjklASDFGHJKL"
	line3 := "zxcvbnmZXCVBNM"
	m := make(map[byte]int, 0)
	// m1, m2, m3 := make(map[byte]int), make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(line1); i++ {
		m[line1[i]] = 1
	}
	for i := 0; i < len(line2); i++ {
		m[line2[i]] = 2
	}
	for i := 0; i < len(line3); i++ {
		m[line3[i]] = 3
	}
	result := make([]string, 0)
	for _, word := range words {
		if valid(m, word) {
			result = append(result, word)
		}
	}
	return result
}

func valid(m map[byte]int, word string) bool {
	first := m[word[0]]
	for i := 1; i < len(word); i++ {
		if m[word[i]] != first {
			return false
		}
	}
	return true
}

