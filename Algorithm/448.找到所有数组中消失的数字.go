/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 *
 * https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/description/
 *
 * algorithms
 * Easy (49.54%)
 * Likes:    170
 * Dislikes: 0
 * Total Accepted:    9K
 * Total Submissions: 18K
 * Testcase Example:  '[4,3,2,7,8,2,3,1]'
 *
 * 给定一个范围在  1 ≤ a[i] ≤ n ( n = 数组大小 ) 的 整型数组，数组中的元素一些出现了两次，另一些只出现一次。
 *
 * 找到所有在 [1, n] 范围之间没有出现在数组中的数字。
 *
 * 您能在不使用额外空间且时间复杂度为O(n)的情况下完成这个任务吗? 你可以假定返回的数组不算在额外空间内。
 *
 * 示例:
 *
 *
 * 输入:
 * [4,3,2,7,8,2,3,1]
 *
 * 输出:
 * [5,6]
 *
 *
 */
func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	if n <= 0 {
		return []int{}
	}
	i := 0
	for i < n {
		if nums[i] == (i + 1) {
			i++
		} else {
			val := nums[i]
			if nums[val-1] == nums[i] {
				i++
			} else {
				nums[val-1], nums[i] = nums[i], nums[val-1]
			}
		}
	}
	result := make([]int, 0)
	for i := 0; i < n; i++ {
		if nums[i] != (i + 1) {
			result = append(result, i+1)
		}
	}
	return result
}

// 4,3,2,7,8,2,3,1
// 7,3,2,4,8,2,3,1
// 3,3,2,4,8,2,7,1
// 2,3,3,4,8,2,7,1
// 3,2,3,4,8,2,7,1


