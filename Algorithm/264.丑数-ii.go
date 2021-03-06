/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 *
 * https://leetcode-cn.com/problems/ugly-number-ii/description/
 *
 * algorithms
 * Medium (44.17%)
 * Likes:    87
 * Dislikes: 0
 * Total Accepted:    5K
 * Total Submissions: 11.3K
 * Testcase Example:  '10'
 *
 * 编写一个程序，找出第 n 个丑数。
 *
 * 丑数就是只包含质因数 2, 3, 5 的正整数。
 *
 * 示例:
 *
 * 输入: n = 10
 * 输出: 12
 * 解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
 *
 * 说明:
 *
 *
 * 1 是丑数。
 * n 不超过1690。
 *
 *
 */
func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}
	result := []int{1}
	index2, index3, index5 := 0, 0, 0
	for len(result) < n {
		val := min(min(result[index2]*2, result[index3]*3), result[index5]*5)
		if val%5 == 0 {
			index5++
		}
		if val%3 == 0 {
			index3++
		}
		if val%2 == 0 {
			index2++
		}
		result = append(result, val)
	}
	return result[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

