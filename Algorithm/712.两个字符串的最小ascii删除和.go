/*
 * @lc app=leetcode.cn id=712 lang=golang
 *
 * [712] 两个字符串的最小ASCII删除和
 *
 * https://leetcode-cn.com/problems/minimum-ascii-delete-sum-for-two-strings/description/
 *
 * algorithms
 * Medium (54.65%)
 * Likes:    45
 * Dislikes: 0
 * Total Accepted:    1.4K
 * Total Submissions: 2.5K
 * Testcase Example:  '"sea"\n"eat"'
 *
 * 给定两个字符串s1, s2，找到使两个字符串相等所需删除字符的ASCII值的最小和。
 *
 * 示例 1:
 *
 *
 * 输入: s1 = "sea", s2 = "eat"
 * 输出: 231
 * 解释: 在 "sea" 中删除 "s" 并将 "s" 的值(115)加入总和。
 * 在 "eat" 中删除 "t" 并将 116 加入总和。
 * 结束时，两个字符串相等，115 + 116 = 231 就是符合条件的最小和。
 *
 *
 * 示例 2:
 *
 *
 * 输入: s1 = "delete", s2 = "leet"
 * 输出: 403
 * 解释: 在 "delete" 中删除 "dee" 字符串变成 "let"，
 * 将 100[d]+101[e]+101[e] 加入总和。在 "leet" 中删除 "e" 将 101[e] 加入总和。
 * 结束时，两个字符串都等于 "let"，结果即为 100+101+101+101 = 403 。
 * 如果改为将两个字符串转换为 "lee" 或 "eet"，我们会得到 433 或 417 的结果，比答案更大。
 *
 *
 * 注意:
 *
 *
 * 0 < s1.length, s2.length <= 1000。
 * 所有字符串中的字符ASCII值在[97, 122]之间。
 *
 *
 */
func minimumDeleteSum(s1 string, s2 string) int {
	m := len(s1)
	n := len(s2)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		dp[i+1][0] = dp[i][0] + int(s1[i])
	}
	for j := 0; j < n; j++ {
		dp[0][j+1] = dp[0][j] + int(s2[j])
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			letter1 := s1[i-1]
			letter2 := s2[j-1]
			if letter1 == letter2 {
				dp[i][j] = dp[i-1][j-1]
			} else {
				val1 := dp[i-1][j-1] + int(letter1) + int(letter2)
				val2 := dp[i-1][j] + int(letter1)
				val3 := dp[i][j-1] + int(letter2)
				dp[i][j] = min(val1, min(val2, val3))
			}
		}
	}
	return dp[m][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


