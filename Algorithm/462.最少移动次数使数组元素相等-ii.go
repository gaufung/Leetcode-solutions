import "sort"

/*
 * @lc app=leetcode.cn id=462 lang=golang
 *
 * [462] 最少移动次数使数组元素相等 II
 *
 * https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements-ii/description/
 *
 * algorithms
 * Medium (52.64%)
 * Likes:    31
 * Dislikes: 0
 * Total Accepted:    1.5K
 * Total Submissions: 2.9K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个非空整数数组，找到使所有数组元素相等所需的最小移动数，其中每次移动可将选定的一个元素加1或减1。 您可以假设数组的长度最多为10000。
 *
 * 例如:
 *
 *
 * 输入:
 * [1,2,3]
 *
 * 输出:
 * 2
 *
 * 说明：
 * 只有两个动作是必要的（记得每一步仅可使其中一个元素加1或减1）：
 *
 * [1,2,3]  =>  [2,2,3]  =>  [2,2,2]
 *
 *
 */
func minMoves2(nums []int) int {
	sort.Ints(nums)
	pivot := nums[len(nums)/2]
	result := 0
	for i := 0; i < len(nums); i++ {
		result += abs(nums[i] - pivot)
	}
	return result
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

