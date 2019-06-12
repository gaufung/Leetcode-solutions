/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 *
 * https://leetcode-cn.com/problems/add-strings/description/
 *
 * algorithms
 * Easy (45.48%)
 * Likes:    73
 * Dislikes: 0
 * Total Accepted:    8.1K
 * Total Submissions: 17.8K
 * Testcase Example:  '"0"\n"0"'
 *
 * 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
 *
 * 注意：
 *
 *
 * num1 和num2 的长度都小于 5100.
 * num1 和num2 都只包含数字 0-9.
 * num1 和num2 都不包含任何前导零。
 * 你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。
 *
 *
 */
func addStrings(num1 string, num2 string) string {
	digit2byte := map[int]byte{
		1: '1',
		2: '2',
		3: '3',
		4: '4',
		5: '5',
		6: '6',
		7: '7',
		8: '8',
		9: '9',
		0: '0',
	}
	byte2digit := map[byte]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}
	carrier := 0
	m, n := len(num1), len(num2)
	i, j := m-1, n-1
	sum := make([]byte, 0)
	for i >= 0 && j >= 0 {
		val := byte2digit[num1[i]] + byte2digit[num2[j]] + carrier
		sum = append(sum, digit2byte[val%10])
		carrier = val / 10
		i--
		j--
	}
	for i >= 0 {
		val := byte2digit[num1[i]] + carrier
		sum = append(sum, digit2byte[val%10])
		carrier = val / 10
		i--
	}
	for j >= 0 {
		val := byte2digit[num2[j]] + carrier
		sum = append(sum, digit2byte[val%10])
		carrier = val / 10
		j--
	}
	if carrier != 0 {
		sum = append(sum, digit2byte[carrier%10])
	}
	for k, l := 0, len(sum)-1; k < l; {
		sum[k], sum[l] = sum[l], sum[k]
		k++
		l--
	}
	return string(sum)
}

