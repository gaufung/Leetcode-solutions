/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	half := sum / 2
	n := len(nums)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, half+1)
	}
	for i := 0; i < half+1; i++ {
		dp[0][i] = nums[0] == i
	}
	for i := 1; i < n; i++ {
		for j := 0; j < half+1; j++ {
			dp[i][j] = dp[i-1][j]
			if nums[i] <= j {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			}
		}
	}
	return dp[n-1][half]
}



