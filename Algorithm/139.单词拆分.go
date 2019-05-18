/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */
func wordBreak(s string, wordDict []string) bool {
	m := make(map[string]bool, 0)
	for _, word := range wordDict {
		m[word] = true
	}
	n := len(s)
	if n == 0 {
		return m[""]
	}
	dp := make([]bool, n+1)
	dp[0] = true
	dp[1] = m[s[:1]]
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && m[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
