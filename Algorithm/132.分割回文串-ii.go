/*
 * @lc app=leetcode.cn id=132 lang=golang
 *
 * [132] 分割回文串 II
 *
 * https://leetcode-cn.com/problems/palindrome-partitioning-ii/description/
 *
 * algorithms
 * Hard (37.67%)
 * Likes:    50
 * Dislikes: 0
 * Total Accepted:    2.3K
 * Total Submissions: 6.2K
 * Testcase Example:  '"aab"'
 *
 * 给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
 *
 * 返回符合要求的最少分割次数。
 *
 * 示例:
 *
 * 输入: "aab"
 * 输出: 1
 * 解释: 进行一次分割就可将 s 分割成 ["aa","b"] 这样两个回文子串。
 *
 *
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


