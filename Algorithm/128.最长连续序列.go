/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 *
 * https://leetcode-cn.com/problems/longest-consecutive-sequence/description/
 *
 * algorithms
 * Hard (43.47%)
 * Likes:    103
 * Dislikes: 0
 * Total Accepted:    7.6K
 * Total Submissions: 17.5K
 * Testcase Example:  '[100,4,200,1,3,2]'
 *
 * 给定一个未排序的整数数组，找出最长连续序列的长度。
 *
 * 要求算法的时间复杂度为 O(n)。
 *
 * 示例:
 *
 * 输入: [100, 4, 200, 1, 3, 2]
 * 输出: 4
 * 解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。
 *
 */
func longestConsecutive(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}
	m := make(map[int]int, 0)
	for i := 0; i < n; i++ {
		m[nums[i]] = i
	}
	max := 0
	for i := 0; i < n; i++ {
		long := travel(m, nums[i])
		if long > max {
			max = long
		}
	}
	return max
}

func travel(m map[int]int, seed int) int {
	if _, ok := m[seed]; ok {
		delete(m, seed)
		return 1 + travel(m, seed-1) + travel(m, seed+1)
	} else {
		return 0
	}
}

