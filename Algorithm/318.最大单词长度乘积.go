/*
 * @lc app=leetcode.cn id=318 lang=golang
 *
 * [318] 最大单词长度乘积
 *
 * https://leetcode-cn.com/problems/maximum-product-of-word-lengths/description/
 *
 * algorithms
 * Medium (56.16%)
 * Likes:    40
 * Dislikes: 0
 * Total Accepted:    1.8K
 * Total Submissions: 3.2K
 * Testcase Example:  '["abcw","baz","foo","bar","xtfn","abcdef"]'
 *
 * 给定一个字符串数组 words，找到 length(word[i]) * length(word[j])
 * 的最大值，并且这两个单词不含有公共字母。你可以认为每个单词只包含小写字母。如果不存在这样的两个单词，返回 0。
 *
 * 示例 1:
 *
 * 输入: ["abcw","baz","foo","bar","xtfn","abcdef"]
 * 输出: 16
 * 解释: 这两个单词为 "abcw", "xtfn"。
 *
 * 示例 2:
 *
 * 输入: ["a","ab","abc","d","cd","bcd","abcd"]
 * 输出: 4
 * 解释: 这两个单词为 "ab", "cd"。
 *
 * 示例 3:
 *
 * 输入: ["a","aa","aaa","aaaa"]
 * 输出: 0
 * 解释: 不存在这样的两个单词。
 *
 */
func maxProduct(words []string) int {
	n := len(words)
	bits := make([]int, n)
	for i := 0; i < n; i++ {
		bits[i] = convert(words[i])
	}
	longest := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if bits[i]&bits[j] == 0 {
				val := len(words[i]) * len(words[j])
				if val > longest {
					longest = val
				}
			}
		}
	}
	return longest
}

func convert(word string) int {
	res := 0
	for i := 0; i < len(word); i++ {
		res |= (0x01 << uint(word[i]-'a'))
	}
	return res
}

