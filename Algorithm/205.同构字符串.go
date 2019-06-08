/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串
 *
 * https://leetcode-cn.com/problems/isomorphic-strings/description/
 *
 * algorithms
 * Easy (44.26%)
 * Likes:    104
 * Dislikes: 0
 * Total Accepted:    11.5K
 * Total Submissions: 25.9K
 * Testcase Example:  '"egg"\n"add"'
 *
 * 给定两个字符串 s 和 t，判断它们是否是同构的。
 *
 * 如果 s 中的字符可以被替换得到 t ，那么这两个字符串是同构的。
 *
 * 所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。
 *
 * 示例 1:
 *
 * 输入: s = "egg", t = "add"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: s = "foo", t = "bar"
 * 输出: false
 *
 * 示例 3:
 *
 * 输入: s = "paper", t = "title"
 * 输出: true
 *
 * 说明:
 * 你可以假设 s 和 t 具有相同的长度。
 *
 */
func isIsomorphic(s string, t string) bool {
	n := len(s)
	m := make(map[byte]byte, 0)
	for i := 0; i < n; i++ {
		if val, ok := m[s[i]]; !ok {
			m[s[i]] = t[i]
		} else {
			if val != t[i] {
				return false
			}
		}
	}
	m = make(map[byte]byte, 0)
	for i := 0; i < n; i++ {
		if val, ok := m[t[i]]; !ok {
			m[t[i]] = s[i]
		} else {
			if val != s[i] {
				return false
			}
		}
	}
	return true
}

