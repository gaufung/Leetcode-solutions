/*
 * @lc app=leetcode.cn id=673 lang=golang
 *
 * [673] 最长递增子序列的个数
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

