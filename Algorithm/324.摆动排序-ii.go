import "sort"

/*
 * @lc app=leetcode.cn id=324 lang=golang
 *
 * [324] 摆动排序 II
 *
 * https://leetcode-cn.com/problems/wiggle-sort-ii/description/
 *
 * algorithms
 * Medium (31.13%)
 * Likes:    37
 * Dislikes: 0
 * Total Accepted:    2.3K
 * Total Submissions: 7.5K
 * Testcase Example:  '[1,5,1,1,6,4]'
 *
 * 给定一个无序的数组 nums，将它重新排列成 nums[0] < nums[1] > nums[2] < nums[3]... 的顺序。
 *
 * 示例 1:
 *
 * 输入: nums = [1, 5, 1, 1, 6, 4]
 * 输出: 一个可能的答案是 [1, 4, 1, 5, 1, 6]
 *
 * 示例 2:
 *
 * 输入: nums = [1, 3, 2, 2, 3, 1]
 * 输出: 一个可能的答案是 [2, 3, 1, 3, 1, 2]
 *
 * 说明:
 * 你可以假设所有输入都会得到有效的结果。
 *
 * 进阶:
 * 你能用 O(n) 时间复杂度和 / 或原地 O(1) 额外空间来实现吗？
 *
 */
func wiggleSort(nums []int) {
	sort.Ints(nums)
	n := len(nums)
	i := n/2 - 1
	j := n - 1
	backup := make([]int, n)
	if n%2 != 0 {
		backup[n-1] = nums[n/2]
	}
	cnt := 0
	for i >= 0 {
		backup[2*cnt] = nums[i]
		backup[2*cnt+1] = nums[j]
		i--
		j--
		cnt++
	}
	for k := 0; k < n; k++ {
		nums[k] = backup[k]
	}

}

