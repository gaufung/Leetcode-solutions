import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=241 lang=golang
 *
 * [241] 为运算表达式设计优先级
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


