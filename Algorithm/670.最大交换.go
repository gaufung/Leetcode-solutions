/*
 * @lc app=leetcode.cn id=670 lang=golang
 *
 * [670] 最大交换
 *
 * https://leetcode-cn.com/problems/maximum-swap/description/
 *
 * algorithms
 * Medium (35.49%)
 * Likes:    30
 * Dislikes: 0
 * Total Accepted:    1.5K
 * Total Submissions: 4.3K
 * Testcase Example:  '2736'
 *
 * 给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。
 *
 * 示例 1 :
 *
 *
 * 输入: 2736
 * 输出: 7236
 * 解释: 交换数字2和数字7。
 *
 *
 * 示例 2 :
 *
 *
 * 输入: 9973
 * 输出: 9973
 * 解释: 不需要交换。
 *
 *
 * 注意:
 *
 *
 * 给定数字的范围是 [0, 10^8]
 *
 *
 */
func maximumSwap(num int) int {
	if num == 0 {
		return 0
	}
	digits := split(num)
	// [6, 3, 7, 2]
	n := len(digits)
	for i := n - 1; i > 0; i-- {
		val := digits[i]
		index := i
		for j := i - 1; j >= 0; j-- {
			if digits[j] >= val {
				index = j
				val = digits[j]
			}
		}
		if index != i && val != digits[i] {
			digits[index], digits[i] = digits[i], digits[index]
			break
		}
	}
	return combine(digits)

}

func combine(digits []int) int {
	val := 0
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		val = val*10 + digits[i]
	}
	return val
}

func split(num int) []int {
	result := make([]int, 0)
	for num > 0 {
		result = append(result, num%10)
		num /= 10
	}
	return result
}

