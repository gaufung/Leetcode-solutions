import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=241 lang=golang
 *
 * [241] 为运算表达式设计优先级
 *
 * https://leetcode-cn.com/problems/different-ways-to-add-parentheses/description/
 *
 * algorithms
 * Medium (67.49%)
 * Likes:    48
 * Dislikes: 0
 * Total Accepted:    1.9K
 * Total Submissions: 2.9K
 * Testcase Example:  '"2-1-1"'
 *
 * 给定一个含有数字和运算符的字符串，为表达式添加括号，改变其运算优先级以求出不同的结果。你需要给出所有可能的组合的结果。有效的运算符号包含 +, - 以及
 * * 。
 *
 * 示例 1:
 *
 * 输入: "2-1-1"
 * 输出: [0, 2]
 * 解释:
 * ((2-1)-1) = 0
 * (2-(1-1)) = 2
 *
 * 示例 2:
 *
 * 输入: "2*3-4*5"
 * 输出: [-34, -14, -10, -10, 10]
 * 解释:
 * (2*(3-(4*5))) = -34
 * ((2*3)-(4*5)) = -14
 * ((2*(3-4))*5) = -10
 * (2*((3-4)*5)) = -10
 * (((2*3)-4)*5) = 10
 *
 */
func diffWaysToCompute(input string) []int {
	operatorCnt := operatorCount(input)
	if operatorCnt == 0 {
		val, _ := strconv.Atoi(input)
		return []int{val}
	}
	if operatorCnt == 1 {
		val := calculateValue(input)
		return []int{val}
	}
	result := make([]int, 0)
	for i := 0; i < len(input); i++ {
		if input[i] == '+' {
			lefts := diffWaysToCompute(input[:i])
			rights := diffWaysToCompute(input[i+1:])
			for _, left := range lefts {
				for _, right := range rights {
					result = append(result, left+right)
				}
			}
		}
		if input[i] == '-' {
			lefts := diffWaysToCompute(input[:i])
			rights := diffWaysToCompute(input[i+1:])
			for _, left := range lefts {
				for _, right := range rights {
					result = append(result, left-right)
				}
			}
		}
		if input[i] == '*' {
			lefts := diffWaysToCompute(input[:i])
			rights := diffWaysToCompute(input[i+1:])
			for _, left := range lefts {
				for _, right := range rights {
					result = append(result, left*right)
				}
			}
		}
	}
	return result
}

func operatorCount(input string) int {
	count := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '+' || input[i] == '-' || input[i] == '*' {
			count++
		}
	}
	return count
}

func calculateValue(input string) int {
	index := strings.IndexByte(input, '+')
	if index == -1 {
		index = strings.IndexByte(input, '-')
	}
	if index == -1 {
		index = strings.IndexByte(input, '*')
	}
	if input[index] == '+' {
		val1, _ := strconv.Atoi(input[:index])
		val2, _ := strconv.Atoi(input[index+1:])
		return val1 + val2
	}
	if input[index] == '-' {
		val1, _ := strconv.Atoi(input[:index])
		val2, _ := strconv.Atoi(input[index+1:])
		return val1 - val2
	}

	val1, _ := strconv.Atoi(input[:index])
	val2, _ := strconv.Atoi(input[index+1:])
	return val1 * val2

}

