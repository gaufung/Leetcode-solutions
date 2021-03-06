/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 *
 * https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/description/
 *
 * algorithms
 * Easy (41.43%)
 * Likes:    146
 * Dislikes: 0
 * Total Accepted:    33.3K
 * Total Submissions: 80.3K
 * Testcase Example:  '[1,2,2,1]\n[2,2]'
 *
 * 给定两个数组，编写一个函数来计算它们的交集。
 *
 * 示例 1:
 *
 * 输入: nums1 = [1,2,2,1], nums2 = [2,2]
 * 输出: [2,2]
 *
 *
 * 示例 2:
 *
 * 输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * 输出: [4,9]
 *
 * 说明：
 *
 *
 * 输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
 * 我们可以不考虑输出结果的顺序。
 *
 *
 * 进阶:
 *
 *
 * 如果给定的数组已经排好序呢？你将如何优化你的算法？
 * 如果 nums1 的大小比 nums2 小很多，哪种方法更优？
 * 如果 nums2 的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
 *
 *
 */
func intersect(nums1 []int, nums2 []int) []int {
	m1, m2 := frequency(nums1), frequency(nums2)
	result := make([]int, 0)
	for key1, val1 := range m1 {
		if val2, ok := m2[key1]; ok {
			for i := 0; i < min(val1, val2); i++ {
				result = append(result, key1)
			}
		}
	}
	return result

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func frequency(nums []int) map[int]int {
	m := make(map[int]int, 0)
	for _, num := range nums {
		if val, ok := m[num]; ok {
			m[num] = val + 1
		} else {
			m[num] = 1
		}
	}
	return m
}


