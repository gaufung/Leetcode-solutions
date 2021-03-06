/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 *
 * https://leetcode-cn.com/problems/first-unique-character-in-a-string/description/
 *
 * algorithms
 * Easy (38.16%)
 * Likes:    117
 * Dislikes: 0
 * Total Accepted:    29.6K
 * Total Submissions: 77.5K
 * Testcase Example:  '"leetcode"'
 *
 * 给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
 *
 * 案例:
 *
 *
 * s = "leetcode"
 * 返回 0.
 *
 * s = "loveleetcode",
 * 返回 2.
 *
 *
 *
 *
 * 注意事项：您可以假定该字符串只包含小写字母。
 *
 */
func firstUniqChar(s string) int {
	m := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		if val, ok := m[s[i]]; ok {
			m[s[i]] = val + 1
		} else {
			m[s[i]] = 1
		}
	}
	for i := 0; i < len(s); i++ {
		if val := m[s[i]]; val == 1 {
			return i
		}
	}
	return -1
}

