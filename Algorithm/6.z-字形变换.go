/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 *
 * https://leetcode-cn.com/problems/zigzag-conversion/description/
 *
 * algorithms
 * Medium (42.52%)
 * Likes:    295
 * Dislikes: 0
 * Total Accepted:    33.3K
 * Total Submissions: 78.3K
 * Testcase Example:  '"PAYPALISHIRING"\n3'
 *
 * 将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。
 *
 * 比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：
 *
 * L   C   I   R
 * E T O E S I I G
 * E   D   H   N
 *
 *
 * 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。
 *
 * 请你实现这个将字符串进行指定行数变换的函数：
 *
 * string convert(string s, int numRows);
 *
 * 示例 1:
 *
 * 输入: s = "LEETCODEISHIRING", numRows = 3
 * 输出: "LCIRETOESIIGEDHN"
 *
 *
 * 示例 2:
 *
 * 输入: s = "LEETCODEISHIRING", numRows = 4
 * 输出: "LDREOEIIECIHNTSG"
 * 解释:
 *
 * L     D     R
 * E   O E   I I
 * E C   I H   N
 * T     S     G
 *
 */
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	} else {
		result := make([]byte, 0)
		i := 0
		for i < numRows {
			if gap, ok := gaps(i, numRows); ok {
				idx := i
				for idx < len(s) {
					result = append(result, s[idx])
					j := idx + gap
					if j < len(s) {
						result = append(result, s[j])
					}
					idx += (2*numRows - 2)
				}
			} else {
				idx := i
				for idx < len(s) {
					result = append(result, s[idx])
					idx += (2*numRows - 2)
				}
			}
			i++
		}
		return string(result)
	}
}

func gaps(i, numRows int) (int, bool) {
	gap := 2*numRows - 2
	if i == 0 || i == numRows-1 {
		return gap, false
	} else {
		return (numRows - 1 - i) * 2, true
	}
}

