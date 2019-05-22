import "strconv"

/*
 * @lc app=leetcode.cn id=228 lang=golang
 *
 * [228] 汇总区间
 *
 * https://leetcode-cn.com/problems/summary-ranges/description/
 *
 * algorithms
 * Medium (47.35%)
 * Likes:    19
 * Dislikes: 0
 * Total Accepted:    2.4K
 * Total Submissions: 5.1K
 * Testcase Example:  '[0,1,2,4,5,7]'
 *
 * 给定一个无重复元素的有序整数数组，返回数组区间范围的汇总。
 *
 * 示例 1:
 *
 * 输入: [0,1,2,4,5,7]
 * 输出: ["0->2","4->5","7"]
 * 解释: 0,1,2 可组成一个连续的区间; 4,5 可组成一个连续的区间。
 *
 * 示例 2:
 *
 * 输入: [0,2,3,4,6,8,9]
 * 输出: ["0","2->4","6","8->9"]
 * 解释: 2,3,4 可组成一个连续的区间; 8,9 可组成一个连续的区间。
 *
 */
func summaryRanges(nums []int) []string {
	result := make([]string, 0)
	n := len(nums)
	if n <= 0 {
		return result
	}
	start := 0
	for start < n {
		end := start
		for end < n-1 {
			if nums[end]+1 == nums[end+1] {
				end++
			} else {
				break
			}
		}
		if start == end {
			result = append(result, strconv.Itoa(nums[start]))
		} else {
			result = append(result, strconv.Itoa(nums[start])+"->"+strconv.Itoa(nums[end]))
		}
		start = end + 1
	}
	return result
}

