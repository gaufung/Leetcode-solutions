import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=537 lang=golang
 *
 * [537] 复数乘法
 *
 * https://leetcode-cn.com/problems/complex-number-multiplication/description/
 *
 * algorithms
 * Medium (64.49%)
 * Likes:    15
 * Dislikes: 0
 * Total Accepted:    2.2K
 * Total Submissions: 3.4K
 * Testcase Example:  '"1+1i"\n"1+1i"'
 *
 * 给定两个表示复数的字符串。
 *
 * 返回表示它们乘积的字符串。注意，根据定义 i^2 = -1 。
 *
 * 示例 1:
 *
 *
 * 输入: "1+1i", "1+1i"
 * 输出: "0+2i"
 * 解释: (1 + i) * (1 + i) = 1 + i^2 + 2 * i = 2i ，你需要将它转换为 0+2i 的形式。
 *
 *
 * 示例 2:
 *
 *
 * 输入: "1+-1i", "1+-1i"
 * 输出: "0+-2i"
 * 解释: (1 - i) * (1 - i) = 1 + i^2 - 2 * i = -2i ，你需要将它转换为 0+-2i 的形式。
 *
 *
 * 注意:
 *
 *
 * 输入字符串不包含额外的空格。
 * 输入字符串将以 a+bi 的形式给出，其中整数 a 和 b 的范围均在 [-100, 100] 之间。输出也应当符合这种形式。
 *
 *
 */
func complexNumberMultiply(a string, b string) string {
	real1, image1 := split(a)
	real2, image2 := split(b)
	real := real1*real2 - image1*image2
	image := real1*image2 + real2*image1
	return fmt.Sprintf("%d+%di", real, image)
}

func split(input string) (int, int) {
	idx := strings.Index(input, "+")
	real, _ := strconv.Atoi(input[:idx])
	image, _ := strconv.Atoi(input[idx+1 : len(input)-1])
	return real, image
}

