/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */
func combine(n int, k int) [][]int {
	return comb(1, n, k)
}

func comb(from, to int, k int) [][]int {
	if k == 0 {
		return [][]int{}
	}
	result := make([][]int, 0)
	if k == 1 {
		for i := from; i <= to; i++ {
			result = append(result, []int{i})
		}
		return result
	}
	for i := from; i <= to-k+1; i++ {
		leftResult := comb(i+1, to, k-1)
		for _, item := range leftResult {
			result = append(result, append(item, i))
		}
	}
	return result
}
