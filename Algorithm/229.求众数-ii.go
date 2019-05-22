/*
 * @lc app=leetcode.cn id=229 lang=golang
 *
 * [229] 求众数 II
 *
 * https://leetcode-cn.com/problems/majority-element-ii/description/
 *
 * algorithms
 * Medium (40.32%)
 * Likes:    68
 * Dislikes: 0
 * Total Accepted:    4.2K
 * Total Submissions: 10.3K
 * Testcase Example:  '[3,2,3]'
 *
 * 给定一个大小为 n 的数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
 *
 * 说明: 要求算法的时间复杂度为 O(n)，空间复杂度为 O(1)。
 *
 * 示例 1:
 *
 * 输入: [3,2,3]
 * 输出: [3]
 *
 * 示例 2:
 *
 * 输入: [1,1,1,3,3,2,2,2]
 * 输出: [1,2]
 *
 */
func majorityElement(nums []int) []int {
	xCnt, yCnt := 0, 0
	x, y := 0, 0
	n := len(nums)
	for _, num := range nums {
		if (xCnt == 0 || num == x) && y != num {
			xCnt++
			x = num
		} else if yCnt == 0 || num == y {
			yCnt++
			y = num
		} else {
			xCnt--
			yCnt--
		}
	}
	result := make([]int, 0)
	cntX := 0
	cntY := 0
	for _, num := range nums {
		if num == x {
			cntX++
		}
		if num == y {
			cntY++
		}
	}
	if cntX > n/3 {
		result = append(result, x)
	}
	if cntY > n/3 {
		if x != y {
			result = append(result, y)
		}

	}
	return result

}

