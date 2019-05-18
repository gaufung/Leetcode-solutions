/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */
func searchInsert(nums []int, target int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}
	lo, hi := 0, n
	for lo < hi {
		mi := (lo + hi) >> 1
		if nums[mi] == target {
			return mi
		} else if nums[mi] > target {
			hi = mi
		} else {
			lo = mi + 1
		}
	}
	return lo
}

