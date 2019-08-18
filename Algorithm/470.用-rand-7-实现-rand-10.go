/*
 * @lc app=leetcode.cn id=470 lang=golang
 *
 * [470] 用 Rand7() 实现 Rand10()
 *
 * https://leetcode-cn.com/problems/implement-rand10-using-rand7/description/
 *
 * algorithms
 * Medium (43.65%)
 * Likes:    26
 * Dislikes: 0
 * Total Accepted:    2.1K
 * Total Submissions: 4.8K
 * Testcase Example:  '1'
 *
 * 已有方法 rand7 可生成 1 到 7 范围内的均匀随机整数，试写一个方法 rand10 生成 1 到 10 范围内的均匀随机整数。
 *
 * 不要使用系统的 Math.random() 方法。
 *
 *
 *
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: 1
 * 输出: [7]
 *
 *
 * 示例 2:
 *
 *
 * 输入: 2
 * 输出: [8,4]
 *
 *
 * 示例 3:
 *
 *
 * 输入: 3
 * 输出: [8,1,10]
 *
 *
 *
 *
 * 提示:
 *
 *
 * rand7 已定义。
 * 传入参数: n 表示 rand10 的调用次数。
 *
 *
 *
 *
 * 进阶:
 *
 *
 * rand7()调用次数的 期望值 是多少 ?
 * 你能否尽量少调用 rand7() ?
 *
 *
 */

/*
 a/b 1 2 3 4 5 6 7
 1   2 3 4 5 6 7 8
 2   3 4 5 6 7 8 9
 3   4 5 6 7 8 9 0
 4   5 6 7 8 9 0 1
 5   6 7 8 9 0 1 2
 6   7 8 9 0 1 2 3
 7   8 9 0 1 2 3 4

 digit frequence
 0 5
 1 4
 2 4
 3 4
 5 4
 6 6
 7 6
 8 7
 9 6
*/

func rand10() int {
	a, b := rand7(), rand7()
	if a >= 5 && b <= 3 {
		return rand10()
	}
	return (a+b)%10 + 1
}

