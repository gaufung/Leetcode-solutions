/*
 * @lc app=leetcode.cn id=373 lang=golang
 *
 * [373] 查找和最小的K对数字
 *
 * https://leetcode-cn.com/problems/find-k-pairs-with-smallest-sums/description/
 *
 * algorithms
 * Medium (35.97%)
 * Likes:    30
 * Dislikes: 0
 * Total Accepted:    1.3K
 * Total Submissions: 3.5K
 * Testcase Example:  '[1,7,11]\n[2,4,6]\n3'
 *
 * 给定两个以升序排列的整形数组 nums1 和 nums2, 以及一个整数 k。
 *
 * 定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2。
 *
 * 找到和最小的 k 对数字 (u1,v1), (u2,v2) ... (uk,vk)。
 *
 * 示例 1:
 *
 * 输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
 * 输出: [1,2],[1,4],[1,6]
 * 解释: 返回序列中的前 3 对数：
 * ⁠    [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
 *
 *
 * 示例 2:
 *
 * 输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
 * 输出: [1,1],[1,1]
 * 解释: 返回序列中的前 2 对数：
 * [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
 *
 *
 * 示例 3:
 *
 * 输入: nums1 = [1,2], nums2 = [3], k = 3
 * 输出: [1,3],[2,3]
 * 解释: 也可能序列中所有的数对都被返回:[1,3],[2,3]
 *
 *
 */
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	result := make([][]int, 0)
	if len(nums1) == 0 || len(nums2) == 0 {
		return result
	}
	i, j := 0, 0
	for k > 0 {
		if i == len(nums1) {
			i = len(nums1) - 1
		}
		if j == len(nums2) {
			j = len(nums2) - 1
		}
		result = append(result, []int{nums1[i], nums2[j]})
		if i == len(nums1)-1 {
			j++
			k--
			continue
		}
		if j == len(nums2)-1 {
			i++
			k--
			continue
		}
		if nums1[i]+nums2[j+1] < nums1[i+1]+nums2[j] {
			j++
		} else if nums1[i]+nums2[j+1] > nums1[i+1]+nums2[j] {
			i++
		} else {
			i++
		}
		k--
	}
	return result
}

