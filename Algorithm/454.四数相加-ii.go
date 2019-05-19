/*
 * @lc app=leetcode.cn id=454 lang=golang
 *
 * [454] 四数相加 II
 *
 * https://leetcode-cn.com/problems/4sum-ii/description/
 *
 * algorithms
 * Medium (51.61%)
 * Likes:    61
 * Dislikes: 0
 * Total Accepted:    5.5K
 * Total Submissions: 10.6K
 * Testcase Example:  '[1,2]\n[-2,-1]\n[-1,2]\n[0,2]'
 *
 * 给定四个包含整数的数组列表 A , B , C , D ,计算有多少个元组 (i, j, k, l) ，使得 A[i] + B[j] + C[k] +
 * D[l] = 0。
 *
 * 为了使问题简单化，所有的 A, B, C, D 具有相同的长度 N，且 0 ≤ N ≤ 500 。所有整数的范围在 -2^28 到 2^28 - 1
 * 之间，最终结果不会超过 2^31 - 1 。
 *
 * 例如:
 *
 *
 * 输入:
 * A = [ 1, 2]
 * B = [-2,-1]
 * C = [-1, 2]
 * D = [ 0, 2]
 *
 * 输出:
 * 2
 *
 * 解释:
 * 两个元组如下:
 * 1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
 * 2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0
 *
 *
 */
func fourSumCount(A []int, B []int, C []int, D []int) int {
	N := len(A)
	if N <= 0 {
		return 0
	}
	m1 := make(map[int]int, 0)
	m2 := make(map[int]int, 0)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			sum1 := A[i] + B[j]
			if val, ok := m1[sum1]; ok {
				m1[sum1] = val + 1
			} else {
				m1[sum1] = 1
			}
			sum2 := C[i] + D[j]
			if val, ok := m2[sum2]; ok {
				m2[sum2] = val + 1
			} else {
				m2[sum2] = 1
			}
		}
	}
	result := 0
	for k1, v1 := range m1 {
		leftk := 0 - k1
		if v2, ok := m2[leftk]; ok {
			result += v1 * v2
		}
	}
	return result
}

