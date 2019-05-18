/*
 * @lc app=leetcode.cn id=740 lang=golang
 *
 * [740] 删除与获得点数
 */
func deleteAndEarn(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}
	array := make([]int, 10001)
	for _, num := range nums {
		array[num] = array[num] + num
	}
	robbery := make([]int, 10001)
	norobbery := make([]int, 10001)
	robbery[0] = 0
	norobbery[0] = 0
	for i := 1; i < 10001; i++ {
		robbery[i] = norobbery[i-1] + array[i]
		norobbery[i] = max(norobbery[i-1], robbery[i-1])
	}
	return max(robbery[10000], norobbery[10000])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

