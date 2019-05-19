/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 *
 * https://leetcode-cn.com/problems/add-binary/description/
 *
 * algorithms
 * Easy (48.17%)
 * Likes:    193
 * Dislikes: 0
 * Total Accepted:    25K
 * Total Submissions: 51.9K
 * Testcase Example:  '"11"\n"1"'
 *
 * 给定两个二进制字符串，返回他们的和（用二进制表示）。
 *
 * 输入为非空字符串且只包含数字 1 和 0。
 *
 * 示例 1:
 *
 * 输入: a = "11", b = "1"
 * 输出: "100"
 *
 * 示例 2:
 *
 * 输入: a = "1010", b = "1011"
 * 输出: "10101"
 *
 */
func addBinary(a string, b string) string {
	m, n := len(a), len(b)
	if m < n {
		return addBinary(b, a)
	}
	carrier := 0
	i, j := m-1, n-1
	result := make([]byte, 0)
	for i >= 0 && j >= 0 {
		add1, add2 := int(a[i]-'0'), int(b[j]-'0')
		digit := add1 ^ add2 ^ carrier
		carrier = (add1 & add2) | (add1 & carrier) | (add2 & carrier)
		if digit == 0 {
			result = append(result, '0')
		} else {
			result = append(result, '1')
		}
		i--
		j--
	}
	for i >= 0 {
		add1, add2 := int(a[i]-'0'), 0
		digit := add1 ^ add2 ^ carrier
		carrier = (add1 & add2) | (add1 & carrier) | (add2 & carrier)
		if digit == 0 {
			result = append(result, '0')
		} else {
			result = append(result, '1')
		}
		i--
	}
	for j >= 0 {
		add1, add2 := 0, int(b[j]-'0')
		digit := add1 ^ add2 ^ carrier
		carrier = (add1 & add2) | (add1 & carrier) | (add2 & carrier)
		if digit == 0 {
			result = append(result, '0')
		} else {
			result = append(result, '1')
		}
		j--
	}
	if carrier == 1 {
		result = append(result, '1')
	}
	for i, j := 0, len(result)-1; i < j; {
		result[i], result[j] = result[j], result[i]
		i++
		j--
	}
	return string(result)
}

