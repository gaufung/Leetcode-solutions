/*
 * @lc app=leetcode.cn id=301 lang=golang
 *
 * [301] 删除无效的括号
 *
 * https://leetcode-cn.com/problems/remove-invalid-parentheses/description/
 *
 * algorithms
 * Hard (39.31%)
 * Likes:    28
 * Dislikes: 0
 * Total Accepted:    1.1K
 * Total Submissions: 2.9K
 * Testcase Example:  '"()())()"'
 *
 * 删除最小数量的无效括号，使得输入的字符串有效，返回所有可能的结果。
 *
 * 说明: 输入可能包含了除 ( 和 ) 以外的字符。
 *
 * 示例 1:
 *
 * 输入: "()())()"
 * 输出: ["()()()", "(())()"]
 *
 *
 * 示例 2:
 *
 * 输入: "(a)())()"
 * 输出: ["(a)()()", "(a())()"]
 *
 *
 * 示例 3:
 *
 * 输入: ")("
 * 输出: [""]
 *
 */
func removeInvalidParentheses(s string) []string {
	leftRemove, rightRemove := needToRemove(s)
	return dfs(s, 0, leftRemove, rightRemove)
}

func dfs(s string, index int, leftRemove, rightRemove int) []string {
	if leftRemove == 0 && rightRemove == 0 && valid(s) {
		return []string{s}
	}
	result := make([]string, 0)
	for i := index; i < len(s); i++ {
		if s[i] == '(' && leftRemove > 0 {
			if i > 0 && s[i-1] == '(' {
				continue
			}
			leftResult := dfs(removeKth(s, i), i, leftRemove-1, rightRemove)
			for _, left := range leftResult {
				result = append(result, left)
			}
		}
		if s[i] == ')' && rightRemove > 0 {
			if i > 0 && s[i-1] == ')' {
				continue
			}
			rightResult := dfs(removeKth(s, i), i, leftRemove, rightRemove-1)
			for _, right := range rightResult {
				result = append(result, right)
			}
		}
	}
	return result
}

func removeKth(s string, i int) string {
	return s[:i] + s[i+1:]
}

func needToRemove(s string) (int, int) {
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else if s[i] == ')' {
			if left > 0 {
				left--
			} else {
				right++
			}
		}
	}
	return left, right
}

func valid(s string) bool {
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		if right > left {
			return false
		}
		if s[i] == '(' {
			left++
		}
		if s[i] == ')' {
			right++
		}
	}
	return left == right
}
