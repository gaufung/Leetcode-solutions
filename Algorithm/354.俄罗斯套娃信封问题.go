import "sort"

/*
 * @lc app=leetcode.cn id=354 lang=golang
 *
 * [354] 俄罗斯套娃信封问题
 *
 * https://leetcode-cn.com/problems/russian-doll-envelopes/description/
 *
 * algorithms
 * Hard (31.53%)
 * Likes:    42
 * Dislikes: 0
 * Total Accepted:    2.2K
 * Total Submissions: 7K
 * Testcase Example:  '[[5,4],[6,4],[6,7],[2,3]]'
 *
 * 给定一些标记了宽度和高度的信封，宽度和高度以整数对形式 (w, h)
 * 出现。当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。
 *
 * 请计算最多能有多少个信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。
 *
 * 说明:
 * 不允许旋转信封。
 *
 * 示例:
 *
 * 输入: envelopes = [[5,4],[6,4],[6,7],[2,3]]
 * 输出: 3
 * 解释: 最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
 *
 *
 */
func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}
	enves := matrix(envelopes)
	sort.Sort(enves)
	result := make([]int, len(enves))
	result[0] = 1
	for i := 1; i < len(enves); i++ {
		result[i] = 1
		for j := i - 1; j >= 0; j-- {
			if isMax(enves[i][0], enves[j][0]) && isMax(enves[i][1], enves[j][1]) {
				result[i] = max(result[j]+1, result[i])
			}
		}
	}
	maxValue := result[0]
	for i := 1; i < len(result); i++ {
		maxValue = max(maxValue, result[i])
	}
	return maxValue
}

func isMax(a, b int) bool {
	return a > b
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

type matrix [][]int

func (m matrix) Len() int { return len(m) }
func (m matrix) Less(i, j int) bool {
	if m[i][0] < m[j][0] {
		return true
	} else if m[i][0] > m[j][0] {
		return false
	} else {
		if m[i][1] < m[j][1] {
			return true
		} else if m[i][1] > m[j][1] {
			return false
		} else {
			return true
		}
	}
}

func (m matrix) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

