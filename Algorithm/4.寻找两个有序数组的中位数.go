/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个有序数组的中位数
 *
 * https://leetcode-cn.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (34.88%)
 * Likes:    1058
 * Dislikes: 0
 * Total Accepted:    55K
 * Total Submissions: 157.8K
 * Testcase Example:  '[1,3]\n[2]'
 *
 * 给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
 *
 * 请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
 *
 * 你可以假设 nums1 和 nums2 不会同时为空。
 *
 * 示例 1:
 *
 * nums1 = [1, 3]
 * nums2 = [2]
 *
 * 则中位数是 2.0
 *
 *
 * 示例 2:
 *
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 *
 * 则中位数是 (2 + 3)/2 = 2.5
 *
 *
 */
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if (m+n)%2 == 0 {
		val1 := findKth(nums1, nums2, (m+n)/2)
		val2 := findKth(nums1, nums2, (m+n)/2+1)
		return (float64(val1) + float64(val2)) * 0.5
	} else {
		val := findKth(nums1, nums2, (m+n+1)/2)
		return float64(val)
	}
}

func findKth(nums1 []int, nums2 []int, k int) int {
	if len(nums1) == 0 {
		return nums2[k-1]
	}
	if len(nums2) == 0 {
		return nums1[k-1]
	}
	if len(nums1) > len(nums2) {
		return findKth(nums2, nums1, k)
	}
	if k == 1 {
		if nums1[0] < nums2[0] {
			return nums1[0]
		} else {
			return nums2[0]
		}
	}
	mi := k >> 1
	mi1, mi2 := mi-1, mi-1
	if mi1 >= len(nums1) {
		mi1 = len(nums1) - 1
	}
	if nums1[mi1] <= nums2[mi2] {
		nums1 := nums1[mi1+1:]
		k = k - (mi1 + 1)
		return findKth(nums1, nums2, k)
	} else {
		nums2 := nums2[mi2+1:]
		k = k - (mi2 + 1)
		return findKth(nums1, nums2, k)
	}
}

