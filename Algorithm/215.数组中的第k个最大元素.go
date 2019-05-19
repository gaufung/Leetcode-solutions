/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 *
 * https://leetcode-cn.com/problems/kth-largest-element-in-an-array/description/
 *
 * algorithms
 * Medium (56.91%)
 * Likes:    170
 * Dislikes: 0
 * Total Accepted:    25.1K
 * Total Submissions: 44K
 * Testcase Example:  '[3,2,1,5,6,4]\n2'
 *
 * 在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
 *
 * 示例 1:
 *
 * 输入: [3,2,1,5,6,4] 和 k = 2
 * 输出: 5
 *
 *
 * 示例 2:
 *
 * 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
 * 输出: 4
 *
 * 说明:
 *
 * 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
 *
 */
func findKthLargest(nums []int, k int) int {

	pivot := nums[0]
	i, j := 0, len(nums)-1
	for i < j {
		for i < j && nums[j] < pivot {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
		for i < j && nums[i] >= pivot {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i] = pivot
	if i+1 == k {
		return nums[i]
	}
	if i+1 > k {
		return findKthLargest(nums[:i], k)
	} else {
		return findKthLargest(nums[i+1:], k-(i+1))
	}
}

