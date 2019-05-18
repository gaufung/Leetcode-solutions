import "fmt"

/*
 * @lc app=leetcode.cn id=698 lang=golang
 *
 * [698] 划分为k个相等的子集
 */
func canPartitionKSubsets(nums []int, k int) bool {
	n := len(nums)
	if n <= 0 {
		return false
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, sum+1)
	}
	for i := 0; i < sum+1; i++ {
		dp[0][i] = nums[0] == i
	}
	for i := 1; i < n; i++ {
		for j := 0; j < sum+1; j++ {
			dp[i][j] = dp[i-1][j]
			if nums[i] <= j {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			}
		}
	}
	result := true
	count := sum / k
	for count <= sum {
		result = result && dp[n-1][count]
		fmt.Printf("%v\n", dp[:][count])
		count += count
	}
	return result
}

