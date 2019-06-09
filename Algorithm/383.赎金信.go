/*
 * @lc app=leetcode.cn id=383 lang=golang
 *
 * [383] 赎金信
 *
 * https://leetcode-cn.com/problems/ransom-note/description/
 *
 * algorithms
 * Easy (47.96%)
 * Likes:    41
 * Dislikes: 0
 * Total Accepted:    7.4K
 * Total Submissions: 15.5K
 * Testcase Example:  '"a"\n"b"'
 *
 * 给定一个赎金信 (ransom)
 * 字符串和一个杂志(magazine)字符串，判断第一个字符串ransom能不能由第二个字符串magazines里面的字符构成。如果可以构成，返回
 * true ；否则返回 false。
 *
 * (题目说明：为了不暴露赎金信字迹，要从杂志上搜索各个需要的字母，组成单词来表达意思。)
 *
 * 注意：
 *
 * 你可以假设两个字符串均只含有小写字母。
 *
 *
 * canConstruct("a", "b") -> false
 * canConstruct("aa", "ab") -> false
 * canConstruct("aa", "aab") -> true
 *
 *
 */
func canConstruct(ransomNote string, magazine string) bool {
	m1 := make(map[byte]int, 0)
	for i := 0; i < len(magazine); i++ {
		val, _ := m1[magazine[i]]
		m1[magazine[i]] = val + 1
	}
	m2 := make(map[byte]int, 0)
	for i := 0; i < len(ransomNote); i++ {
		val, _ := m2[ransomNote[i]]
		m2[ransomNote[i]] = val + 1
	}
	for k1, v1 := range m2 {
		if m1[k1] < v1 {
			return false
		}
	}
	return true
}

