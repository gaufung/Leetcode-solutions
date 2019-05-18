import "sort"

/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 *
 * https://leetcode-cn.com/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (35.68%)
 * Likes:    117
 * Dislikes: 0
 * Total Accepted:    14.7K
 * Total Submissions: 41.1K
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * 给出一个区间的集合，请合并所有重叠的区间。
 *
 * 示例 1:
 *
 * 输入: [[1,3],[2,6],[8,10],[15,18]]
 * 输出: [[1,6],[8,10],[15,18]]
 * 解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
 *
 *
 * 示例 2:
 *
 * 输入: [[1,4],[4,5]]
 * 输出: [[1,5]]
 * 解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
 *
 */
func merge(intervals [][]int) [][]int {
	n := len(intervals)
	components := make([]Component, n*2)
	for i, interval := range intervals {
		components[i*2] = Component{0, interval[0]}
		components[i*2+1] = Component{1, interval[1]}
	}
	sort.Slice(components, func(i, j int) bool {
		if components[i].Val < components[j].Val {
			return true
		}
		if components[i].Val > components[j].Val {
			return false
		}
		return components[i].Type < components[j].Type
	})
	result := make([][]int, 0)
	stack := make([]Component, 0)
	for _, c := range components {
		if c.Type == 0 {
			stack = append(stack, c)
		} else {
			if len(stack) == 1 {
				result = append(result, []int{stack[0].Val, c.Val})
			}
			stack = stack[:len(stack)-1]
		}
	}
	return result
}

type Component struct {
	Type int
	Val  int
}

