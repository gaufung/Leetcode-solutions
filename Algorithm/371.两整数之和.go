/*
 * @lc app=leetcode.cn id=371 lang=golang
 *
 * [371] 两整数之和
 *
 * https://leetcode-cn.com/problems/sum-of-two-integers/description/
 *
 * algorithms
 * Easy (52.01%)
 * Likes:    107
 * Dislikes: 0
 * Total Accepted:    10.1K
 * Total Submissions: 19.4K
 * Testcase Example:  '1\n2'
 *
 * 不使用运算符 + 和 - ​​​​​​​，计算两整数 ​​​​​​​a 、b ​​​​​​​之和。
 *
 * 示例 1:
 *
 * 输入: a = 1, b = 2
 * 输出: 3
 *
 *
 * 示例 2:
 *
 * 输入: a = -2, b = 3
 * 输出: 1
 *
 */
func getSum(a int, b int) int {
	mask := uint32(0x01)
	val := uint32(0x00)
	carrier := uint32(0x00)
	newA, newB := uint32(a), uint32(b)
	for i := uint(0); i < 32; i++ {
		add1 := (newA >> i) & mask
		add2 := (newB >> i) & mask
		digit := add1 ^ add2 ^ carrier
		carrier = (add1 & add2) | (add1 & carrier) | (add2 & carrier)
		val |= digit << i
	}
	return int(int32(val))
}

