/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 *
 * https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/description/
 *
 * algorithms
 * Easy (37.41%)
 * Likes:    122
 * Dislikes: 0
 * Total Accepted:    5.8K
 * Total Submissions: 15.3K
 * Testcase Example:  '"cbaebabacd"\n"abc"'
 *
 * 给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。
 *
 * 字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。
 *
 * 说明：
 *
 *
 * 字母异位词指字母相同，但排列不同的字符串。
 * 不考虑答案输出的顺序。
 *
 *
 * 示例 1:
 *
 *
 * 输入:
 * s: "cbaebabacd" p: "abc"
 *
 * 输出:
 * [0, 6]
 *
 * 解释:
 * 起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
 * 起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。
 *
 *
 * 示例 2:
 *
 *
 * 输入:
 * s: "abab" p: "ab"
 *
 * 输出:
 * [0, 1, 2]
 *
 * 解释:
 * 起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
 * 起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
 * 起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词。
 *
 *
 */
func findAnagrams(s string, p string) []int {
	n, m := len(s), len(p)
	if n < m {
		return []int{}
	}
	result := make([]int, 0)
	m2 := make(map[byte]int, 0)
	for i := 0; i < m; i++ {
		val := m2[p[i]]
		m2[p[i]] = val + 1
	}
	m1 := make(map[byte]int, 0)
	for i := 0; i < m-1; i++ {
		val := m1[s[i]]
		m1[s[i]] = val + 1
	}
	for i := 0; i <= n-m; i++ {
		if i > 0 {
			val := m1[s[i-1]]
			if val == 1 {
				delete(m1, s[i-1])
			} else {
				m1[s[i-1]] = val - 1
			}
		}

		val := m1[s[i+m-1]]
		m1[s[i+m-1]] = val + 1

		if xor(m1, m2) {
			result = append(result, i)
		}
	}
	return result
}

func xor(m1, m2 map[byte]int) bool {
	for k1, v1 := range m1 {
		if v2, ok := m2[k1]; ok {
			if v1 != v2 {
				return false
			}
		} else {
			return false
		}
	}
	for k2, v2 := range m2 {
		if v1, ok := m1[k2]; ok {
			if v1 != v2 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

