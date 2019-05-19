import "sort"

/*
 * @lc app=leetcode.cn id=368 lang=golang
 *
 * [368] 最大整除子集
 *
 * https://leetcode-cn.com/problems/largest-divisible-subset/description/
 *
 * algorithms
 * Medium (32.85%)
 * Likes:    32
 * Dislikes: 0
 * Total Accepted:    1.5K
 * Total Submissions: 4.4K
 * Testcase Example:  '[1,2,3]'
 *
 * 给出一个由无重复的正整数组成的集合，找出其中最大的整除子集，子集中任意一对 (Si，Sj) 都要满足：Si % Sj = 0 或 Sj % Si =
 * 0。
 *
 * 如果有多个目标子集，返回其中任何一个均可。
 *
 *
 *
 * 示例 1:
 *
 * 输入: [1,2,3]
 * 输出: [1,2] (当然, [1,3] 也正确)
 *
 *
 * 示例 2:
 *
 * 输入: [1,2,4,8]
 * 输出: [1,2,4,8]
 *
 *
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

