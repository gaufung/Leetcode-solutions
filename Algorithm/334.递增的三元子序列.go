/*
 * @lc app=leetcode.cn id=334 lang=golang
 *
 * [334] 递增的三元子序列
 *
 * https://leetcode-cn.com/problems/increasing-triplet-subsequence/description/
 *
 * algorithms
 * Medium (33.91%)
 * Likes:    55
 * Dislikes: 0
 * Total Accepted:    6.1K
 * Total Submissions: 17.9K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 给定一个未排序的数组，判断这个数组中是否存在长度为 3 的递增子序列。
 *
 * 数学表达式如下:
 *
 * 如果存在这样的 i, j, k,  且满足 0 ≤ i < j < k ≤ n-1，
 * 使得 arr[i] < arr[j] < arr[k] ，返回 true ; 否则返回 false 。
 *
 * 说明: 要求算法的时间复杂度为 O(n)，空间复杂度为 O(1) 。
 *
 * 示例 1:
 *
 * 输入: [1,2,3,4,5]
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: [5,4,3,2,1]
 * 输出: false
 *
 */
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	mins := make([]int, len(nums))
	maxes := make([]int, len(nums))
	mins[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < mins[i-1] {
			mins[i] = nums[i]
		} else {
			mins[i] = mins[i-1]
		}
	}
	maxes[len(nums)-1] = nums[len(nums)-1]
	for j := len(nums) - 2; j >= 0; j-- {
		if nums[j] > maxes[j+1] {
			maxes[j] = nums[j]
		} else {
			maxes[j] = maxes[j+1]
		}
	}
	for i := 0; i < len(nums); i++ {
		if mins[i] < nums[i] && nums[i] < maxes[i] {
			return true
		}
	}
	return false
}

