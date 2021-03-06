import "strconv"

/*
 * @lc app=leetcode.cn id=60 lang=golang
 *
 * [60] 第k个排列
 *
 * https://leetcode-cn.com/problems/permutation-sequence/description/
 *
 * algorithms
 * Medium (46.01%)
 * Likes:    79
 * Dislikes: 0
 * Total Accepted:    8.2K
 * Total Submissions: 17.8K
 * Testcase Example:  '3\n3'
 *
 * 给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。
 *
 * 按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：
 *
 *
 * "123"
 * "132"
 * "213"
 * "231"
 * "312"
 * "321"
 *
 *
 * 给定 n 和 k，返回第 k 个排列。
 *
 * 说明：
 *
 *
 * 给定 n 的范围是 [1, 9]。
 * 给定 k 的范围是[1,  n!]。
 *
 *
 * 示例 1:
 *
 * 输入: n = 3, k = 3
 * 输出: "213"
 *
 *
 * 示例 2:
 *
 * 输入: n = 4, k = 9
 * 输出: "2314"
 *
 *
 */
func getPermutation(n int, k int) string {
	nums := make([]string, n)
	for i := 0; i < n; i++ {
		nums[i] = strconv.Itoa(i + 1)
	}
	return permutaion(nums, k, "")
}

func permutaion(nums []string, k int, prefix string) string {
	if k == 0 {
		return prefix
	} else if len(nums) == 1 {
		return prefix + nums[0]
	} else {
		factor := factorial(len(nums) - 1)
		idx := 0
		for k > factor {
			idx++
			k = k - factor
		}
		newNums := make([]string, 0)
		for i := 0; i < len(nums); i++ {
			if i != idx {
				newNums = append(newNums, nums[i])
			}
		}
		return permutaion(newNums, k, prefix+nums[idx])

	}
}

func factorial(n int) int {
	if n <= 1 {
		return n
	} else {
		return n * factorial(n-1)
	}
}

