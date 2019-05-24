import (
	"sort"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=386 lang=golang
 *
 * [386] 字典序排数
 *
 * https://leetcode-cn.com/problems/lexicographical-numbers/description/
 *
 * algorithms
 * Medium (58.31%)
 * Likes:    26
 * Dislikes: 0
 * Total Accepted:    1.5K
 * Total Submissions: 2.5K
 * Testcase Example:  '13'
 *
 * 给定一个整数 n, 返回从 1 到 n 的字典顺序。
 *
 * 例如，
 *
 * 给定 n =1 3，返回 [1,10,11,12,13,2,3,4,5,6,7,8,9] 。
 *
 * 请尽可能的优化算法的时间复杂度和空间复杂度。 输入的数据 n 小于等于 5,000,000。
 *
 */
func lexicalOrder(n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = i + 1
	}
	sort.Slice(result, func(i, j int) bool {
		return strconv.Itoa(result[i]) < strconv.Itoa(result[j])
	})
	return result
}

