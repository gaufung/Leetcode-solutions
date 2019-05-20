/*
 * @lc app=leetcode.cn id=168 lang=golang
 *
 * [168] Excel表列名称
 *
 * https://leetcode-cn.com/problems/excel-sheet-column-title/description/
 *
 * algorithms
 * Easy (32.25%)
 * Likes:    110
 * Dislikes: 0
 * Total Accepted:    8.1K
 * Total Submissions: 25.2K
 * Testcase Example:  '1'
 *
 * 给定一个正整数，返回它在 Excel 表中相对应的列名称。
 *
 * 例如，
 *
 * ⁠   1 -> A
 * ⁠   2 -> B
 * ⁠   3 -> C
 * ⁠   ...
 * ⁠   26 -> Z
 * ⁠   27 -> AA
 * ⁠   28 -> AB
 * ⁠   ...
 *
 *
 * 示例 1:
 *
 * 输入: 1
 * 输出: "A"
 *
 *
 * 示例 2:
 *
 * 输入: 28
 * 输出: "AB"
 *
 *
 * 示例 3:
 *
 * 输入: 701
 * 输出: "ZY"
 *
 *
 */
func convertToTitle(n int) string {
	// str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// m := make(map[int]byte, 0)
	// for i := 0; i < len(str); i++ {
	// 	m[i+1] = str[i]
	// }
	// result := make([]byte, 0)
	// for n > 26 {
	// 	val := n % 26
	// 	result = append(result, m[val])
	// 	n = n / 26
	// }
	// result = append(result, m[n])
	// for i, j := 0, len(result)-1; i < j; {
	// 	result[i], result[j] = result[j], result[i]
	// 	i++
	// 	j--
	// }
	// return string(result)
	result := make([]byte, 0)
	for n > 0 {
		n--
		result = append(result, byte('A'+n%26))
		n /= 26
	}
	for i, j := 0, len(result)-1; i < j; {
		result[i], result[j] = result[j], result[i]
		i++
		j--
	}
	return string(result)
}

