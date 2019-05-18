import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=368 lang=golang
 *
 * [368] 最大整除子集
 */
func largestDivisibleSubset(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}
	if n == 1 {
		return []int{nums[0]}
	}
	sort.Ints(nums)
	dp := make([]int, n)
	dp[0] = 1
	previous := make([]int, n)
	previous[0] = -1
	maxLength := 1
	for i := 1; i < n; i++ {
		val := 0
		prev := -1
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 {
				if dp[j] > val {
					prev = j
					val = dp[j]
				}
			}
		}
		dp[i] = val + 1
		previous[i] = prev
		maxLength = max(maxLength, dp[i])
	}
	result := make([]int, 0)
	i := 0
	for {
		if dp[i] == maxLength {
			break
		}
		i++
	}
	for i >= 0 {
		result = append(result, nums[i])
		i = previous[i]
	}
	return result

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
