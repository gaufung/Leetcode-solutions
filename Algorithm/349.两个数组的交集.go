import "sort"

/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 *
 * https://leetcode-cn.com/problems/intersection-of-two-arrays/description/
 *
 * algorithms
 * Easy (63.11%)
 * Likes:    98
 * Dislikes: 0
 * Total Accepted:    22.2K
 * Total Submissions: 35.2K
 * Testcase Example:  '[1,2,2,1]\n[2,2]'
 *
 * 给定两个数组，编写一个函数来计算它们的交集。
 *
 * 示例 1:
 *
 * 输入: nums1 = [1,2,2,1], nums2 = [2,2]
 * 输出: [2]
 *
 *
 * 示例 2:
 *
 * 输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * 输出: [9,4]
 *
 * 说明:
 *
 *
 * 输出结果中的每个元素一定是唯一的。
 * 我们可以不考虑输出结果的顺序。
 *
 *
 */
func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	result := make([]int, 0)
	n, m := len(nums1), len(nums2)
	i, j := 0, 0
	for i < n && j < m {
		if nums1[i] == nums2[j] {
			result = append(result, nums1[i])
			i++
			j++
			for i < n && nums1[i] == nums1[i-1] {
				i++
			}
			for j < m && nums2[j] == nums2[j-1] {
				j++
			}
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return result
}

