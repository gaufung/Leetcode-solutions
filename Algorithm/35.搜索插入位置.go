/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 *
 * https://leetcode-cn.com/problems/search-insert-position/description/
 *
 * algorithms
 * Easy (43.37%)
 * Likes:    249
 * Dislikes: 0
 * Total Accepted:    45.7K
 * Total Submissions: 105.4K
 * Testcase Example:  '[1,3,5,6]\n5'
 *
 * 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
 *
 * 你可以假设数组中无重复元素。
 *
 * 示例 1:
 *
 * 输入: [1,3,5,6], 5
 * 输出: 2
 *
 *
 * 示例 2:
 *
 * 输入: [1,3,5,6], 2
 * 输出: 1
 *
 *
 * 示例 3:
 *
 * 输入: [1,3,5,6], 7
 * 输出: 4
 *
 *
 * 示例 4:
 *
 * 输入: [1,3,5,6], 0
 * 输出: 0
 *
 *
 */
func searchInsert(nums []int, target int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}
	lo, hi := 0, n
	for lo < hi {
		mi := (lo + hi) >> 1
		if nums[mi] == target {
			return mi
		} else if nums[mi] > target {
			hi = mi
		} else {
			lo = mi + 1
		}
	}
	return lo
}
