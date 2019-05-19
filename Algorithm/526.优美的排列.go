/*
 * @lc app=leetcode.cn id=526 lang=golang
 *
 * [526] 优美的排列
 *
 * https://leetcode-cn.com/problems/beautiful-arrangement/description/
 *
 * algorithms
 * Medium (51.90%)
 * Likes:    26
 * Dislikes: 0
 * Total Accepted:    1K
 * Total Submissions: 1.9K
 * Testcase Example:  '2'
 *
 * 假设有从 1 到 N 的 N 个整数，如果从这 N 个数字中成功构造出一个数组，使得数组的第 i 位 (1 <= i <= N)
 * 满足如下两个条件中的一个，我们就称这个数组为一个优美的排列。条件：
 *
 *
 * 第 i 位的数字能被 i 整除
 * i 能被第 i 位上的数字整除
 *
 *
 * 现在给定一个整数 N，请问可以构造多少个优美的排列？
 *
 * 示例1:
 *
 *
 * 输入: 2
 * 输出: 2
 * 解释:
 *
 * 第 1 个优美的排列是 [1, 2]:
 * ⁠ 第 1 个位置（i=1）上的数字是1，1能被 i（i=1）整除
 * ⁠ 第 2 个位置（i=2）上的数字是2，2能被 i（i=2）整除
 *
 * 第 2 个优美的排列是 [2, 1]:
 * ⁠ 第 1 个位置（i=1）上的数字是2，2能被 i（i=1）整除
 * ⁠ 第 2 个位置（i=2）上的数字是1，i（i=2）能被 1 整除
 *
 *
 * 说明:
 *
 *
 * N 是一个正整数，并且不会超过15。
 *
 *
 */
func countArrangement(N int) int {
	nums := make([]int, N)
	for i := 0; i < N; i++ {
		nums[i] = i + 1
	}
	return arrange(nums, 1)
}

func arrange(nums []int, startIndex int) int {
	if len(nums) == 1 {
		if nums[0]%startIndex == 0 {
			return 1
		}
		if startIndex%nums[0] == 0 {
			return 1
		}
		return 0
	}
	result := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%startIndex == 0 || startIndex%nums[i] == 0 {
			leftNums := removeKth(nums, i)
			result += arrange(leftNums, startIndex+1)
		}
	}
	return result
}

func removeKth(nums []int, k int) []int {
	result := make([]int, 0)
	for i, val := range nums {
		if i != k {
			result = append(result, val)
		}
	}
	return result
}

