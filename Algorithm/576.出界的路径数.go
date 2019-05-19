/*
 * @lc app=leetcode.cn id=576 lang=golang
 *
 * [576] 出界的路径数
 *
 * https://leetcode-cn.com/problems/out-of-boundary-paths/description/
 *
 * algorithms
 * Medium (32.26%)
 * Likes:    18
 * Dislikes: 0
 * Total Accepted:    707
 * Total Submissions: 2.2K
 * Testcase Example:  '2\n2\n2\n0\n0'
 *
 * 给定一个 m × n 的网格和一个球。球的起始坐标为 (i,j)
 * ，你可以将球移到相邻的单元格内，或者往上、下、左、右四个方向上移动使球穿过网格边界。但是，你最多可以移动 N
 * 次。找出可以将球移出边界的路径数量。答案可能非常大，返回 结果 mod 10^9 + 7 的值。
 *
 *
 *
 * 示例 1：
 *
 * 输入: m = 2, n = 2, N = 2, i = 0, j = 0
 * 输出: 6
 * 解释:
 *
 *
 *
 * 示例 2：
 *
 * 输入: m = 1, n = 3, N = 3, i = 0, j = 1
 * 输出: 12
 * 解释:
 *
 *
 *
 *
 *
 * 说明:
 *
 *
 * 球一旦出界，就不能再被移动回网格内。
 * 网格的长度和高度在 [1,50] 的范围内。
 * N 在 [0,50] 的范围内。
 *
 */
func findPaths(m int, n int, N int, i int, j int) int {
	dp := make([][][]int, N+1)
	for k := 0; k < N+1; k++ {
		dp[k] = buildMap(m, n)
	}
	dp[0][i][j] = 1
	cnt := 0
	for k := 0; k < N; k++ {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if dp[k][i][j] != 0 {
					if i-1 < 0 {
						cnt += dp[k][i][j]
					} else {
						dp[k+1][i-1][j] += dp[k][i][j]
					}
					if i+1 >= m {
						cnt += dp[k][i][j]
					} else {
						dp[k+1][i+1][j] += dp[k][i][j]
					}
					if j-1 < 0 {
						cnt += dp[k][i][j]
					} else {
						dp[k+1][i][j-1] += dp[k][i][j]
					}
					if j+1 >= n {
						cnt += dp[k][i][j]
					} else {
						dp[k+1][i][j+1] += dp[k][i][j]
					}
				}
			}
		}
	}
	threshold := 1000000000 + 7
	if cnt > threshold {
		return cnt % threshold
	}
	return cnt

}

func buildMap(m, n int) [][]int {
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
	}
	return result
}

