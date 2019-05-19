/*
 * @lc app=leetcode.cn id=673 lang=golang
 *
 * [673] 最长递增子序列的个数
 *
 * https://leetcode-cn.com/problems/number-of-longest-increasing-subsequence/description/
 *
 * algorithms
 * Medium (28.92%)
 * Likes:    48
 * Dislikes: 0
 * Total Accepted:    1.2K
 * Total Submissions: 4.2K
 * Testcase Example:  '[1,3,5,4,7]'
 *
 * 给定一个未排序的整数数组，找到最长递增子序列的个数。
 *
 * 示例 1:
 *
 *
 * 输入: [1,3,5,4,7]
 * 输出: 2
 * 解释: 有两个最长递增子序列，分别是 [1, 3, 4, 7] 和[1, 3, 5, 7]。
 *
 *
 * 示例 2:
 *
 *
 * 输入: [2,2,2,2,2]
 * 输出: 5
 * 解释: 最长递增子序列的长度是1，并且存在5个子序列的长度为1，因此输出5。
 *
 *
 * 注意: 给定的数组长度不超过 2000 并且结果一定是32位有符号整数。
 *
 */
func findNumberOfLIS(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}
	longest := make([]int, n)
	longest[0] = 1
	max := 1
	for i := 1; i < n; i++ {
		val := 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] && longest[j]+1 > val {
				val = longest[j] + 1
			}
		}
		longest[i] = val
		if longest[i] > max {
			max = longest[i]
		}
	}
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		if longest[i] != 1 {
			for j := i - 1; j >= 0; j-- {
				if longest[j] == longest[i]-1 && nums[i] > nums[j] {
					dp[i] += dp[j]
				}
			}
		} else {
			dp[i] = 1
		}
	}
	cnt := 0
	for i := 0; i < n; i++ {
		if longest[i] == max {
			cnt += dp[i]
		}
	}
	return cnt
}
