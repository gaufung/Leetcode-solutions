/*
 * @lc app=leetcode.cn id=132 lang=golang
 *
 * [132] 分割回文串 II
 */
func minCut(s string) int {
	n := len(s)
	if n <= 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i < n+1; i++ {
		val := i + 1
		for j := 0; j < i; j++ {
			if valid(s[j:i]) {
				val = min(val, dp[j]+1)
			}
		}
		dp[i] = val
	}
	return dp[n] - 1
}

func valid(s string) bool {
	lo, hi := 0, len(s)-1
	for lo < hi {
		if s[lo] != s[hi] {
			return false
		}
		lo++
		hi--
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

