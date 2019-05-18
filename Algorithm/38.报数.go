/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 报数
 *
 * https://leetcode-cn.com/problems/count-and-say/description/
 *
 * algorithms
 * Easy (49.99%)
 * Likes:    229
 * Dislikes: 0
 * Total Accepted:    30.9K
 * Total Submissions: 61.8K
 * Testcase Example:  '1'
 *
 * 报数序列是一个整数序列，按照其中的整数的顺序进行报数，得到下一个数。其前五项如下：
 *
 * 1.     1
 * 2.     11
 * 3.     21
 * 4.     1211
 * 5.     111221
 *
 *
 * 1 被读作  "one 1"  ("一个一") , 即 11。
 * 11 被读作 "two 1s" ("两个一"）, 即 21。
 * 21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。
 *
 * 给定一个正整数 n（1 ≤ n ≤ 30），输出报数序列的第 n 项。
 *
 * 注意：整数顺序将表示为一个字符串。
 *
 *
 *
 * 示例 1:
 *
 * 输入: 1
 * 输出: "1"
 *
 *
 * 示例 2:
 *
 * 输入: 4
 * 输出: "1211"
 *
 *
 */
func countAndSay(n int) string {
	if n <= 0 {
		return ""
	}
	result := make([]int, 1)
	result[0] = 1
	for n > 1 {
		result = count(result)
		n--
	}
	m := map[int]byte{
		0: '0',
		1: '1',
		2: '2',
		3: '3',
		4: '4',
		5: '5',
		6: '6',
		7: '7',
		8: '8',
		9: '9',
	}
	seq := make([]byte, len(result))
	for i := 0; i < len(result); i++ {
		seq[i] = m[result[i]]
	}
	return string(seq)
}

func count(seq []int) []int {
	first, second := 0, 0
	result := make([]int, 0)
	for second < len(seq) {
		if seq[first] == seq[second] {
			second++
		} else {
			result = append(result, (second - first), seq[first])
			first = second
		}
	}
	result = append(result, (second - first), seq[first])
	return result
}

