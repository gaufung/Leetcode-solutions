/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 *
 * https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (49.03%)
 * Likes:    53
 * Dislikes: 0
 * Total Accepted:    9.3K
 * Total Submissions: 19K
 * Testcase Example:  '[3,4,5,1,2]'
 *
 * 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
 *
 * ( 例如，数组 [0,1,2,4,5,6,7]  可能变为 [4,5,6,7,0,1,2] )。
 *
 * 请找出其中最小的元素。
 *
 * 你可以假设数组中不存在重复元素。
 *
 * 示例 1:
 *
 * 输入: [3,4,5,1,2]
 * 输出: 1
 *
 * 示例 2:
 *
 * 输入: [4,5,6,7,0,1,2]
 * 输出: 0
 *
 */
func findMin(nums []int) int {
	lo, hi := 0, len(nums)
	for lo < hi {
		mi := (lo + hi) >> 1
		if nums[mi] > nums[lo] && nums[mi] > nums[hi-1] {
			lo = mi + 1
		} else {
			if mi != 0 && mi == len(nums)-1 && nums[mi] < nums[mi-1] {
				return nums[mi]
			}
			if mi != 0 && mi != len(nums)-1 && nums[mi] < nums[mi+1] && nums[mi] < nums[mi-1] {
				return nums[mi]
			} else {
				hi = mi
			}
		}
	}
	return nums[lo]
}

