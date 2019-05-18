/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 *
 * https://leetcode-cn.com/problems/sort-colors/description/
 *
 * algorithms
 * Medium (52.15%)
 * Likes:    163
 * Dislikes: 0
 * Total Accepted:    18.9K
 * Total Submissions: 36.2K
 * Testcase Example:  '[2,0,2,1,1,0]'
 *
 * 给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
 *
 * 此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
 *
 * 注意:
 * 不能使用代码库中的排序函数来解决这道题。
 *
 * 示例:
 *
 * 输入: [2,0,2,1,1,0]
 * 输出: [0,0,1,1,2,2]
 *
 * 进阶：
 *
 *
 * 一个直观的解决方案是使用计数排序的两趟扫描算法。
 * 首先，迭代计算出0、1 和 2 元素的个数，然后按照0、1、2的排序，重写当前数组。
 * 你能想出一个仅使用常数空间的一趟扫描算法吗？
 *
 *
 */
func sortColors(nums []int) {
	n := len(nums)
	if n <= 0 {
		return
	}
	zero, two := lastZero(nums, 0), firstTwo(nums, n-1)
	k := zero + 1
	for k < two {
		if nums[k] == 0 {
			nums[zero+1], nums[k] = nums[k], nums[zero+1]
			zero = lastZero(nums, zero+1)
			k = zero + 1
		} else if nums[k] == 2 {
			nums[two-1], nums[k] = nums[k], nums[two-1]
			two = firstTwo(nums, two-1)
		} else {
			k++
		}
	}
}

func lastZero(nums []int, from int) int {
	for from < len(nums) && nums[from] == 0 {
		from++
	}
	return from - 1
}

func firstTwo(nums []int, from int) int {
	for from >= 0 && nums[from] == 2 {
		from--
	}
	return from + 1
}

