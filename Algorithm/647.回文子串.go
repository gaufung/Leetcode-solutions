/*
 * @lc app=leetcode.cn id=647 lang=golang
 *
 * [647] 回文子串
 */
func countSubstrings(s string) int {
	n := len(s)
	if n <= 0 {
		return 0
	}
	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if valid(s[j : i+1]) {
				sum++
			}
		}
	}
	return sum

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
