/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 *
 * https://leetcode-cn.com/problems/first-missing-positive/description/
 *
 * algorithms
 * Hard (35.26%)
 * Likes:    175
 * Dislikes: 0
 * Total Accepted:    13.3K
 * Total Submissions: 37.8K
 * Testcase Example:  '[1,2,0]'
 *
 * 给定一个未排序的整数数组，找出其中没有出现的最小的正整数。
 *
 * 示例 1:
 *
 * 输入: [1,2,0]
 * 输出: 3
 *
 *
 * 示例 2:
 *
 * 输入: [3,4,-1,1]
 * 输出: 2
 *
 *
 * 示例 3:
 *
 * 输入: [7,8,9,11,12]
 * 输出: 1
 *
 *
 * 说明:
 *
 * 你的算法的时间复杂度应为O(n)，并且只能使用常数级别的空间。
 *
 */
func firstMissingPositive(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return 1
	}
	i := 0
	for i < n {
		if nums[i] == i+1 {
			i++
			continue
		}
		if nums[i] > n {
			i++
			continue
		}
		if nums[i] <= 0 {
			i++
			continue
		}
		swapIndex := nums[i] - 1
		if nums[swapIndex] != nums[i] {
			nums[swapIndex], nums[i] = nums[i], nums[swapIndex]
		} else {
			i++
		}
	}
	i = 0
	for ; i < n; i++ {
		if nums[i] != i+1 {
			break
		}
	}
	return i + 1
}

