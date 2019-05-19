/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 *
 * https://leetcode-cn.com/problems/move-zeroes/description/
 *
 * algorithms
 * Easy (54.18%)
 * Likes:    321
 * Dislikes: 0
 * Total Accepted:    46.2K
 * Total Submissions: 85.3K
 * Testcase Example:  '[0,1,0,3,12]'
 *
 * 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
 *
 * 示例:
 *
 * 输入: [0,1,0,3,12]
 * 输出: [1,3,12,0,0]
 *
 * 说明:
 *
 *
 * 必须在原数组上操作，不能拷贝额外的数组。
 * 尽量减少操作次数。
 *
 *
 */
func moveZeroes(nums []int) {
	first, second := 0, 0
	for {
		first = forwardZero(nums, first)
		second = forwardNonZero(nums, first)
		if first < len(nums) && second < len(nums) {
			nums[first], nums[second] = nums[second], nums[first]
			first = first + 1
		} else {
			break
		}
	}

}

func forwardNonZero(nums []int, from int) int {
	for from < len(nums) && nums[from] == 0 {
		from++
	}
	return from
}

func forwardZero(nums []int, from int) int {
	for from < len(nums) && nums[from] != 0 {
		from++
	}
	return from
}

