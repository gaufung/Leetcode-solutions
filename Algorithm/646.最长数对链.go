import "sort"

/*
 * @lc app=leetcode.cn id=646 lang=golang
 *
 * [646] 最长数对链
 *
 * https://leetcode-cn.com/problems/maximum-length-of-pair-chain/description/
 *
 * algorithms
 * Medium (49.05%)
 * Likes:    36
 * Dislikes: 0
 * Total Accepted:    1.6K
 * Total Submissions: 3.3K
 * Testcase Example:  '[[1,2], [2,3], [3,4]]'
 *
 * 给出 n 个数对。 在每一个数对中，第一个数字总是比第二个数字小。
 *
 * 现在，我们定义一种跟随关系，当且仅当 b < c 时，数对(c, d) 才可以跟在 (a, b) 后面。我们用这种形式来构造一个数对链。
 *
 * 给定一个对数集合，找出能够形成的最长数对链的长度。你不需要用到所有的数对，你可以以任何顺序选择其中的一些数对来构造。
 *
 * 示例 :
 *
 *
 * 输入: [[1,2], [2,3], [3,4]]
 * 输出: 2
 * 解释: 最长的数对链是 [1,2] -> [3,4]
 *
 *
 * 注意：
 *
 *
 * 给出数对的个数在 [1, 1000] 范围内。
 *
 *
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


