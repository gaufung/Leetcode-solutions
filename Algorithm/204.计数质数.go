/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 *
 * https://leetcode-cn.com/problems/count-primes/description/
 *
 * algorithms
 * Easy (27.62%)
 * Likes:    154
 * Dislikes: 0
 * Total Accepted:    17.6K
 * Total Submissions: 63.7K
 * Testcase Example:  '10'
 *
 * 统计所有小于非负整数 n 的质数的数量。
 *
 * 示例:
 *
 * 输入: 10
 * 输出: 4
 * 解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
 *
 *
 */
func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	nums := make([]bool, n+1)
	for i := 2; i < n; i++ {
		if nums[i] == false {
			for j := i + i; j < n; j += i {
				nums[j] = true
			}
		}
	}
	cnt := 0
	for i := 2; i < n; i++ {
		if nums[i] == false {
			cnt++
		}
	}
	return cnt
}

