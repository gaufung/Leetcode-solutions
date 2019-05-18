/*
 * @lc app=leetcode.cn id=376 lang=golang
 *
 * [376] 摆动序列
 */
func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}
	ups := make([]int, n)
	downs := make([]int, n)
	ups[0] = 1
	downs[0] = 1
	result := 1
	for i := 1; i < n; i++ {
		up := 0
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				up = max(up, downs[j])
			}
		}
		ups[i] = up + 1
		down := 0
		for j := i - 1; j >= 0; j-- {
			if nums[i] < nums[j] {
				down = max(down, ups[j])
			}
		}
		downs[i] = down + 1
		result = max(result, max(ups[i], downs[i]))
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

