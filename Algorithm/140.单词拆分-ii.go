import "strings"

/*
 * @lc app=leetcode.cn id=140 lang=golang
 *
 * [140] 单词拆分 II
 *
 * https://leetcode-cn.com/problems/word-break-ii/description/
 *
 * algorithms
 * Hard (37.38%)
 * Likes:    45
 * Dislikes: 0
 * Total Accepted:    3.6K
 * Total Submissions: 9.5K
 * Testcase Example:  '"catsanddog"\n["cat","cats","and","sand","dog"]'
 *
 * 给定一个非空字符串 s 和一个包含非空单词列表的字典
 * wordDict，在字符串中增加空格来构建一个句子，使得句子中所有的单词都在词典中。返回所有这些可能的句子。
 *
 * 说明：
 *
 *
 * 分隔时可以重复使用字典中的单词。
 * 你可以假设字典中没有重复的单词。
 *
 *
 * 示例 1：
 *
 * 输入:
 * s = "catsanddog"
 * wordDict = ["cat", "cats", "and", "sand", "dog"]
 * 输出:
 * [
 * "cats and dog",
 * "cat sand dog"
 * ]
 *
 *
 * 示例 2：
 *
 * 输入:
 * s = "pineapplepenapple"
 * wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
 * 输出:
 * [
 * "pine apple pen apple",
 * "pineapple pen apple",
 * "pine applepen apple"
 * ]
 * 解释: 注意你可以重复使用字典中的单词。
 *
 *
 * 示例 3：
 *
 * 输入:
 * s = "catsandog"
 * wordDict = ["cats", "dog", "sand", "and", "cat"]
 * 输出:
 * []
 *
 *
 */
func wordBreak(s string, wordDict []string) []string {
	m := toDictionary(wordDict)
	n := len(s)
	dp := make([][][]string, n+1)
	dp[0] = [][]string{
		[]string{},
	}
	for i := 1; i < n+1; i++ {
		val := [][]string{}
		for j := i - 1; j >= 0; j-- {
			if _, ok := m[s[j:i]]; ok && (j == 0 || len(dp[j]) > 0) {
				str := s[j:i]

				for _, item := range dp[j] {
					vals := make([]string, 0)
					for _, ss := range item {
						vals = append(vals, ss)
					}
					vals = append(vals, str)
					val = append(val, vals)
				}

			}
		}
		dp[i] = val
	}
	result := make([]string, 0)
	for _, items := range dp[n] {
		result = append(result, strings.Join(items, " "))
	}
	return result

}

func toDictionary(wordDict []string) map[string]bool {
	m := make(map[string]bool, 0)
	for _, word := range wordDict {
		m[word] = true
	}
	return m
}

