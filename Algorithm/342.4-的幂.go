/*
 * @lc app=leetcode.cn id=342 lang=golang
 *
 * [342] 4的幂
 *
 * https://leetcode-cn.com/problems/power-of-four/description/
 *
 * algorithms
 * Easy (45.42%)
 * Likes:    57
 * Dislikes: 0
 * Total Accepted:    8.5K
 * Total Submissions: 18.6K
 * Testcase Example:  '16'
 *
 * 给定一个整数 (32 位有符号整数)，请编写一个函数来判断它是否是 4 的幂次方。
 *
 * 示例 1:
 *
 * 输入: 16
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: 5
 * 输出: false
 *
 * 进阶：
 * 你能不使用循环或者递归来完成本题吗？
 *
 */
func isPowerOfFour(num int) bool {
	oneCount := 0
	rightZeroCount := 0
	isFoundOne := false
	mask := 0x01
	for num > 0 {
		if num&mask == 1 {
			oneCount++
			isFoundOne = true
		} else {
			if !isFoundOne {
				rightZeroCount++
			}
		}
		num = num >> 1
	}
	return oneCount == 1 && rightZeroCount%2 == 0
}

