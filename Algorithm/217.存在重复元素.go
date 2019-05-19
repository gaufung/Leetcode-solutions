/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 *
 * https://leetcode-cn.com/problems/contains-duplicate/description/
 *
 * algorithms
 * Easy (48.07%)
 * Likes:    141
 * Dislikes: 0
 * Total Accepted:    56.8K
 * Total Submissions: 118.2K
 * Testcase Example:  '[1,2,3,1]'
 *
 * 给定一个整数数组，判断是否存在重复元素。
 *
 * 如果任何值在数组中出现至少两次，函数返回 true。如果数组中每个元素都不相同，则返回 false。
 *
 * 示例 1:
 *
 * 输入: [1,2,3,1]
 * 输出: true
 *
 * 示例 2:
 *
 * 输入: [1,2,3,4]
 * 输出: false
 *
 * 示例 3:
 *
 * 输入: [1,1,1,3,3,4,3,2,4,2]
 * 输出: true
 *
 */
func containsDuplicate(nums []int) bool {
	m := make(map[int]bool, 0)
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		} else {
			m[num] = true
		}
	}
	return false
}

