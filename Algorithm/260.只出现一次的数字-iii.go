/*
 * @lc app=leetcode.cn id=260 lang=golang
 *
 * [260] 只出现一次的数字 III
 *
 * https://leetcode-cn.com/problems/single-number-iii/description/
 *
 * algorithms
 * Medium (64.00%)
 * Likes:    80
 * Dislikes: 0
 * Total Accepted:    5.1K
 * Total Submissions: 8K
 * Testcase Example:  '[1,2,1,3,2,5]'
 *
 * 给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。
 *
 * 示例 :
 *
 * 输入: [1,2,1,3,2,5]
 * 输出: [3,5]
 *
 * 注意：
 *
 *
 * 结果输出的顺序并不重要，对于上面的例子， [5, 3] 也是正确答案。
 * 你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现？
 *
 *
 */
func singleNumber(nums []int) []int {
	res := 0
	for _, num := range nums {
		res ^= num
	}
	rightShift := uint(0)
	mask := 0x01
	for rightShift < 32 {
		if (res>>rightShift)&mask == mask {
			break
		}
		rightShift++
	}
	val1, val2 := 0, 0
	for _, num := range nums {
		if (num>>rightShift)&mask == mask {
			val1 ^= num
		} else {
			val2 ^= num
		}
	}
	return []int{val1, val2}
}

