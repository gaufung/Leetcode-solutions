/*
 * @lc app=leetcode.cn id=442 lang=golang
 *
 * [442] 数组中重复的数据
 *
 * https://leetcode-cn.com/problems/find-all-duplicates-in-an-array/description/
 *
 * algorithms
 * Medium (60.86%)
 * Likes:    85
 * Dislikes: 0
 * Total Accepted:    4.7K
 * Total Submissions: 7.7K
 * Testcase Example:  '[4,3,2,7,8,2,3,1]'
 *
 * 给定一个整数数组 a，其中1 ≤ a[i] ≤ n （n为数组长度）, 其中有些元素出现两次而其他元素出现一次。
 *
 * 找到所有出现两次的元素。
 *
 * 你可以不用到任何额外空间并在O(n)时间复杂度内解决这个问题吗？
 *
 * 示例：
 *
 *
 * 输入:
 * [4,3,2,7,8,2,3,1]
 *
 * 输出:
 * [2,3]
 *
 *
 */
func findDuplicates(nums []int) []int {
	// [4,3,2,7,8,2,3,1]
	// [7, 3, 2, 4, 8, 2, 3, 1]
	// [3, 3, 2, 4, 8, 2, 7, 1]
	// [2, 3, 3, 4, 8, 2, 7, 1]
	// [3, 2, 3, 4, 8, 2, 7, 1]
	// [3, 2, 3, 4, 1, 2, 7, 8]
	// [1, 2, 3, 4, 3, 2, 7, 8]
	// [2, 3, 3, 2, 4, 1]
	// [3, 2, 3, 2, 4, 1]
	// [3, 2, 3, 2, 4, 1]
	// result := make(map[int]bool, 0)
	// n := len(nums)
	// if n < 2 {
	// 	return []int{}
	// }
	// i := 0
	// for i < n {
	// 	if i+1 == nums[i] {
	// 		i++
	// 	} else {
	// 		if nums[i] == nums[nums[i]-1] {
	// 			//result = append(result, nums[i])
	// 			result[nums[i]] = true
	// 			i++
	// 		} else {
	// 			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
	// 		}
	// 	}
	// }
	// res := make([]int, 0)
	// for k, _ := range result {
	// 	res = append(res, k)
	// }
	// return res
	result := make([]int, 0)
	for _, num := range nums {
		num = abs(num)
		if nums[num-1] > 0 {
			nums[num-1] *= -1
		} else {
			result = append(result, num)
		}
	}
	return result
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

