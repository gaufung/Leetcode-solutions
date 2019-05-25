/*
 * @lc app=leetcode.cn id=389 lang=golang
 *
 * [389] 找不同
 *
 * https://leetcode-cn.com/problems/find-the-difference/description/
 *
 * algorithms
 * Easy (56.49%)
 * Likes:    70
 * Dislikes: 0
 * Total Accepted:    9.5K
 * Total Submissions: 16.8K
 * Testcase Example:  '"abcd"\n"abcde"'
 *
 * 给定两个字符串 s 和 t，它们只包含小写字母。
 *
 * 字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。
 *
 * 请找出在 t 中被添加的字母。
 *
 *
 *
 * 示例:
 *
 * 输入：
 * s = "abcd"
 * t = "abcde"
 *
 * 输出：
 * e
 *
 * 解释：
 * 'e' 是那个被添加的字母。
 *
 *
 */
func findTheDifference(s string, t string) byte {
	res := 0
	for i := 0; i < len(s); i++ {
		res ^= int(s[i] - 'a')
	}
	for i := 0; i < len(t); i++ {
		res ^= int(t[i] - 'a')
	}
	return byte(int('a') + res)
}

