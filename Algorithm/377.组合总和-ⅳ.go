/*
 * @lc app=leetcode.cn id=377 lang=golang
 *
 * [377] 组合总和 Ⅳ
 */
func combinationSum4(nums []int, target int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i < target+1; i++ {
		val := 0
		for _, num := range nums {
			if i >= num {
				val += dp[i-num]
			}
		}
		dp[i] = val
	}
	return dp[target]

}

