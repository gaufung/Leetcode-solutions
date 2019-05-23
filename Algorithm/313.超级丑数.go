/*
 * @lc app=leetcode.cn id=313 lang=golang
 *
 * [313] 超级丑数
 *
 * https://leetcode-cn.com/problems/super-ugly-number/description/
 *
 * algorithms
 * Medium (53.56%)
 * Likes:    25
 * Dislikes: 0
 * Total Accepted:    1.7K
 * Total Submissions: 3.2K
 * Testcase Example:  '12\n[2,7,13,19]'
 *
 * 编写一段程序来查找第 n 个超级丑数。
 *
 * 超级丑数是指其所有质因数都是长度为 k 的质数列表 primes 中的正整数。
 *
 * 示例:
 *
 * 输入: n = 12, primes = [2,7,13,19]
 * 输出: 32
 * 解释: 给定长度为 4 的质数列表 primes = [2,7,13,19]，前 12
 * 个超级丑数序列为：[1,2,4,7,8,13,14,16,19,26,28,32] 。
 *
 * 说明:
 *
 *
 * 1 是任何给定 primes 的超级丑数。
 * 给定 primes 中的数字以升序排列。
 * 0 < k ≤ 100, 0 < n ≤ 10^6, 0 < primes[i] < 1000 。
 * 第 n 个超级丑数确保在 32 位有符整数范围内。
 *
 *
 */
func nthSuperUglyNumber(n int, primes []int) int {
	if n == 1 {
		return 1
	}
	k := len(primes)
	indexes := make([]int, k)
	result := make([]int, n)
	result[0] = 1
	max_int32 := int((^uint(0) >> 1))
	for i := 1; i < n; i++ {
		val := max_int32
		for j := 0; j < k; j++ {
			val = min(val, primes[j]*result[indexes[j]])
		}
		for j := 0; j < k; j++ {
			if val%primes[j] == 0 {
				indexes[j] = indexes[j] + 1
			}
		}
		result[i] = val
	}
	return result[n-1]

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

