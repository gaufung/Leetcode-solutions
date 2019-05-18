import "sort"

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 *
 * https://leetcode-cn.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (33.29%)
 * Likes:    550
 * Dislikes: 0
 * Total Accepted:    80.2K
 * Total Submissions: 240.8K
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * 编写一个函数来查找字符串数组中的最长公共前缀。
 *
 * 如果不存在公共前缀，返回空字符串 ""。
 *
 * 示例 1:
 *
 * 输入: ["flower","flow","flight"]
 * 输出: "fl"
 *
 *
 * 示例 2:
 *
 * 输入: ["dog","racecar","car"]
 * 输出: ""
 * 解释: 输入不存在公共前缀。
 *
 *
 * 说明:
 *
 * 所有输入只包含小写字母 a-z 。
 *
 */
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	strings := Strings(strs)
	sort.Sort(strings)
	first := strings[0]
	last := strings[len(strs)-1]
	i := 0
	for ; i < len(last) && i < len(first); i++ {
		if last[i] != first[i] {
			break
		}
	}
	return string(last[:i])
}

type Strings []string

func (s Strings) Len() int { return len(s) }
func (s Strings) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Strings) Less(i, j int) bool {
	return s[i] < s[j]
}

