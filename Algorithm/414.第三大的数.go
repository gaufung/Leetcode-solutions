/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] 第三大的数
 *
 * https://leetcode-cn.com/problems/third-maximum-number/description/
 *
 * algorithms
 * Easy (31.64%)
 * Likes:    62
 * Dislikes: 0
 * Total Accepted:    8.5K
 * Total Submissions: 26.8K
 * Testcase Example:  '[3,2,1]'
 *
 * 给定一个非空数组，返回此数组中第三大的数。如果不存在，则返回数组中最大的数。要求算法时间复杂度必须是O(n)。
 *
 * 示例 1:
 *
 *
 * 输入: [3, 2, 1]
 *
 * 输出: 1
 *
 * 解释: 第三大的数是 1.
 *
 *
 * 示例 2:
 *
 *
 * 输入: [1, 2]
 *
 * 输出: 2
 *
 * 解释: 第三大的数不存在, 所以返回最大的数 2 .
 *
 *
 * 示例 3:
 *
 *
 * 输入: [2, 2, 3, 1]
 *
 * 输出: 1
 *
 * 解释: 注意，要求返回第三大的数，是指第三大且唯一出现的数。
 * 存在两个值为2的数，它们都排第二。
 *
 *
 */
func thirdMax(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	first, second, third := nums[0], nums[0], nums[0]
	isSecond, isThird := false, false
	for i := 1; i < n; i++ {
		if nums[i] > first {
			tmp := first
			first = nums[i]
			if isSecond {
				third = second
				second = tmp
				isThird = true
			} else {
				second = tmp
				isSecond = true
			}
		} else if nums[i] == first {
			continue
		} else {
			if isSecond {
				if nums[i] > second {
					tmp := second
					second = nums[i]
					third = tmp
					isThird = true
				} else if nums[i] == second {
					continue
				} else {
					if isThird {
						if nums[i] > third {
							third = nums[i]
						}
					} else {
						third = nums[i]
						isThird = true
					}
				}
			} else {
				second = nums[i]
				isSecond = true
			}
		}
	}
	if isThird {
		return third
	} else {
		return first
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

