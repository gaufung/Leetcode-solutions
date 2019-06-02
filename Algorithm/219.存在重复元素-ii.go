/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 *
 * https://leetcode-cn.com/problems/contains-duplicate-ii/description/
 *
 * algorithms
 * Easy (34.43%)
 * Likes:    87
 * Dislikes: 0
 * Total Accepted:    14.3K
 * Total Submissions: 41.5K
 * Testcase Example:  '[1,2,3,1]\n3'
 *
 * 给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j
 * 的差的绝对值最大为 k。
 *
 * 示例 1:
 *
 * 输入: nums = [1,2,3,1], k = 3
 * 输出: true
 *
 * 示例 2:
 *
 * 输入: nums = [1,0,1,1], k = 1
 * 输出: true
 *
 * 示例 3:
 *
 * 输入: nums = [1,2,3,1,2,3], k = 2
 * 输出: false
 *
 */
func containsNearbyDuplicate(nums []int, k int) bool {
	// m := make(map[int][]int, 0)
	// for i := 0; i < len(nums); i++ {
	// 	if val, ok := m[nums[i]]; ok {
	// 		m[nums[i]] = append(val, i)
	// 	} else {
	// 		m[nums[i]] = []int{i}
	// 	}
	// }
	// for _, indexes := range m {
	// 	n := len(indexes)
	// 	if indexes[n-1]-indexes[0] == k {
	// 		return true
	// 	}
	// }
	// return false
	n := len(nums)
	if n < k {
		return false
	}
	for i := 0; i < n-k; i++ {
		if nums[i] == nums[i+k] {
			return true
		}
	}
	return false
}

