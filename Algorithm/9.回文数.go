/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else {
		backup := x
		result := 0
		maxInt := maxInt()
		for backup > 0 {
			digit := backup % 10
			result = result*10 + digit
			backup = backup / 10
			if result > maxInt {
				return false
			}
		}
		return result == x
	}
}
func maxInt() int {
	return int(^uint(0) >> 1)
}
