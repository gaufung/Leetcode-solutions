/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] 插入区间
 *
 * https://leetcode-cn.com/problems/insert-interval/description/
 *
 * algorithms
 * Hard (34.35%)
 * Likes:    42
 * Dislikes: 0
 * Total Accepted:    5.2K
 * Total Submissions: 15.3K
 * Testcase Example:  '[[1,3],[6,9]]\n[2,5]'
 *
 * 给出一个无重叠的 ，按照区间起始端点排序的区间列表。
 *
 * 在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。
 *
 * 示例 1:
 *
 * 输入: intervals = [[1,3],[6,9]], newInterval = [2,5]
 * 输出: [[1,5],[6,9]]
 *
 *
 * 示例 2:
 *
 * 输入: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
 * 输出: [[1,2],[3,10],[12,16]]
 * 解释: 这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。
 *
 *
 */
func insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0)
	n := len(intervals)
	if n <= 0 {
		return [][]int{newInterval}
	}
	if newInterval[1] < intervals[0][0] {
		result = append([][]int{newInterval}, intervals...)
		return result
	}
	if newInterval[1] == intervals[0][0] {
		result = append([][]int{[]int{newInterval[0], intervals[0][1]}}, intervals[1:]...)
		return result
	}
	if newInterval[0] > intervals[n-1][1] {
		result = append(intervals, newInterval)
		return result
	}
	if newInterval[0] == intervals[n-1][1] {
		result = append(intervals[:n-1], []int{intervals[n-1][0], newInterval[1]})
		return result
	}
	i := findLess(intervals, newInterval[0])
	j := findGreat(intervals, newInterval[1])

	insertFrom := min(intervals[i][0], newInterval[0])
	insertTo := max(intervals[j][1], newInterval[1])

	for k := 0; k < i; k++ {
		result = append(result, intervals[k])
	}
	result = append(result, []int{insertFrom, insertTo})
	for k := j + 1; k < len(intervals); k++ {
		result = append(result, intervals[k])
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findGreat(intervals [][]int, target int) int {
	lo, hi := 0, len(intervals)
	for lo < hi {
		mi := (lo + hi) >> 1
		if intervals[mi][0] > target {
			hi = mi
		} else {
			lo = mi + 1
		}
	}
	return lo - 1
}

func findLess(intervals [][]int, target int) int {
	lo, hi := 0, len(intervals)
	for lo < hi {
		mi := (lo + hi) >> 1
		if intervals[mi][1] < target {
			lo = mi + 1
		} else {
			hi = mi
		}
	}
	return lo
}


