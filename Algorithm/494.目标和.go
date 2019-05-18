/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */
func findTargetSumWays(nums []int, S int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}
	dp := make([]map[int]int, n)
	firstM := map[int]int{
		nums[0]: 1,
	}
	val := firstM[-nums[0]]
	firstM[-nums[0]] = val + 1
	dp[0] = firstM
	for i := 1; i < n; i++ {
		m := dp[i-1]
		newM := make(map[int]int, 0)
		for k, v := range m {
			newKey := k + nums[i]
			if val, ok := newM[newKey]; ok {
				newM[newKey] = val + v
			} else {
				newM[newKey] = v
			}
			newKey = k - nums[i]
			if val, ok := newM[newKey]; ok {
				newM[newKey] = val + v
			} else {
				newM[newKey] = v
			}
		}
		dp[i] = newM
	}
	return dp[n-1][S]
}

