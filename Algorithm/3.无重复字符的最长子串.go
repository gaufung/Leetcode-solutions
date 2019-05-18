/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 *
 * https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (29.42%)
 * Likes:    1814
 * Dislikes: 0
 * Total Accepted:    123.5K
 * Total Submissions: 419.5K
 * Testcase Example:  '"abcabcbb"'
 *
 * 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
 *
 * 示例 1:
 *
 * 输入: "abcabcbb"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 *
 *
 * 示例 2:
 *
 * 输入: "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 *
 *
 * 示例 3:
 *
 * 输入: "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 *
 *
 */
func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
func lengthOfLongestSubstring(s string) int {
	result := 0
	maps := make(map[rune]int)
	for idx, char := range s {
		if start, ok := maps[char]; ok {
			result = max(result, len(maps))
			maps = make(map[rune]int)
			for ii, val := range s[start : idx+1] {
				maps[val] = start + ii
			}
		} else {
			maps[char] = idx
		}
	}
	result = max(result, len(maps))
	return result

}

