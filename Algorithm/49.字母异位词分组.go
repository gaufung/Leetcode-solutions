import "sort"

/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 *
 * https://leetcode-cn.com/problems/group-anagrams/description/
 *
 * algorithms
 * Medium (55.80%)
 * Likes:    141
 * Dislikes: 0
 * Total Accepted:    16.3K
 * Total Submissions: 29.2K
 * Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
 *
 * 给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
 *
 * 示例:
 *
 * 输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
 * 输出:
 * [
 * ⁠ ["ate","eat","tea"],
 * ⁠ ["nat","tan"],
 * ⁠ ["bat"]
 * ]
 *
 * 说明：
 *
 *
 * 所有输入均为小写字母。
 * 不考虑答案输出的顺序。
 *
 *
 */
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string, 0)
	for _, str := range strs {
		key := stringToKey(str)
		if val, ok := m[key]; ok {
			m[key] = append(val, str)
		} else {
			m[key] = []string{str}
		}
	}
	result := make([][]string, 0)
	for _, items := range m {
		result = append(result, items)
	}
	return result
}

func stringToKey(str string) string {
	bytes := []byte(str)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	return string(bytes)
}

