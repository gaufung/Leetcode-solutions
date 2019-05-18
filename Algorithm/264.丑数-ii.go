import "fmt"

/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 */
func nthUglyNumber(n int) int {
	if n <= 0 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = 1
	index2, index3, index5 := 0, 0, 0
	for i := 1; i < n; i++ {
		candidate := min(dp[index2]*2, min(dp[index3]*3, dp[index5]*5))
		if candidate%5 == 0 {
			index5++
		}
		if candidate%3 == 0 {
			index3++
		}
		if candidate%2 == 0 {
			index2++
		}
		dp[i] = candidate
	}
	fmt.Printf("%v\n", dp)
	return dp[n-1]

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
