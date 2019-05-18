/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */
func rob(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	withFirst := nums[1:]
	withLast := nums[:n-1]
	withFirstLast := nums[1 : n-1]
	return max(max(robbery(withFirst), robbery(withLast)),
		robbery(withFirstLast))
}

func robbery(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}
	robbery := make([]int, n)
	noRobbery := make([]int, n)
	robbery[0] = nums[0]
	noRobbery[0] = 0
	for i := 1; i < n; i++ {
		robbery[i] = noRobbery[i-1] + nums[i]
		noRobbery[i] = max(noRobbery[i-1], robbery[i-1])
	}
	return max(robbery[n-1], noRobbery[n-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

