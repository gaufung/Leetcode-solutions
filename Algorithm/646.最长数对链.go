import "sort"

/*
 * @lc app=leetcode.cn id=646 lang=golang
 *
 * [646] 最长数对链
 */
func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][0] < pairs[j][0] {
			return true
		} else if pairs[i][0] > pairs[j][0] {
			return false
		} else {
			return pairs[i][1] < pairs[j][1]
		}
	})
	dp := make([]int, len(pairs))
	dp[0] = 1
	max := dp[0]
	for i := 1; i < len(pairs); i++ {
		val := 1
		for j := i - 1; j >= 0; j-- {
			if pairs[i][0] > pairs[j][1] {
				if dp[j]+1 > val {
					val = dp[j] + 1
				}
			}
		}
		dp[i] = val
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

