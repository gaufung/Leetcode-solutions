/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 *
 * https://leetcode-cn.com/problems/decode-ways/description/
 *
 * algorithms
 * Medium (20.39%)
 * Likes:    126
 * Dislikes: 0
 * Total Accepted:    7.6K
 * Total Submissions: 37.5K
 * Testcase Example:  '"12"'
 *
 * 一条包含字母 A-Z 的消息通过以下方式进行了编码：
 *
 * 'A' -> 1
 * 'B' -> 2
 * ...
 * 'Z' -> 26
 *
 *
 * 给定一个只包含数字的非空字符串，请计算解码方法的总数。
 *
 * 示例 1:
 *
 * 输入: "12"
 * 输出: 2
 * 解释: 它可以解码为 "AB"（1 2）或者 "L"（12）。
 *
 *
 * 示例 2:
 *
 * 输入: "226"
 * 输出: 3
 * 解释: 它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。
 *
 *
 */
func numDecodings(s string) int {
	n := len(s)
	dp := make([]int, n)
	if isValid([]byte(s)[:1]) {
		dp[0] = 1
	} else {
		dp[0] = 0
	}
	if n == 1 {
		return dp[0]
	}
	if isValid([]byte(s)[:2]) {
		dp[1] += 1
	}
	if isValid([]byte(s)[:1]) && isValid([]byte(s)[1:2]) {
		dp[1] += 1
	}
	for i := 2; i < n; i++ {
		if isValid([]byte(s)[i : i+1]) {
			dp[i] += dp[i-1]
		}
		if dp[i-2] != 0 && isValid([]byte(s)[i-1:i+1]) {
			dp[i] += dp[i-2]
		}
	}
	return dp[n-1]
}

func isValid(b []byte) bool {
	if len(b) == 0 {
		return true
	}
	if len(b) == 1 {
		return b[0] >= '1' && b[0] <= '9'
	}
	if len(b) == 2 {
		if b[0] == '1' && b[1] >= '0' && b[1] <= '9' {
			return true
		}
		if b[0] == '2' && b[1] >= '0' && b[1] <= '6' {
			return true
		}
		return false
	}
	return true
}

