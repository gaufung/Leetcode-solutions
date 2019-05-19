/*
 * @lc app=leetcode.cn id=421 lang=golang
 *
 * [421] 数组中两个数的最大异或值
 *
 * https://leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array/description/
 *
 * algorithms
 * Medium (54.20%)
 * Likes:    44
 * Dislikes: 0
 * Total Accepted:    1K
 * Total Submissions: 1.9K
 * Testcase Example:  '[3,10,5,25,2,8]'
 *
 * 给定一个非空数组，数组中元素为 a0, a1, a2, … , an-1，其中 0 ≤ ai < 2^31 。
 *
 * 找到 ai 和aj 最大的异或 (XOR) 运算结果，其中0 ≤ i,  j < n 。
 *
 * 你能在O(n)的时间解决这个问题吗？
 *
 * 示例:
 *
 *
 * 输入: [3, 10, 5, 25, 2, 8]
 *
 * 输出: 28
 *
 * 解释: 最大的结果是 5 ^ 25 = 28.
 *
 *
 */
func findMaximumXOR(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}
	root := &TreeNode{Val: -1}
	mask := 0x01
	for _, num := range nums {
		cur := root
		for j := 31; j >= 0; j-- {
			i := uint(j)
			if (num>>i)&mask == 0 {
				if cur.Left == nil {
					cur.Left = &TreeNode{Val: 0}
				}
				cur = cur.Left
			} else {
				if cur.Right == nil {
					cur.Right = &TreeNode{Val: 1}
				}
				cur = cur.Right
			}
		}
		cur.Left = &TreeNode{Val: num}
	}
	result := 0
	for _, num := range nums {
		cur := root
		for j := 31; j >= 0; j-- {
			i := uint(j)
			if (num>>i)&mask == 1 {
				if cur.Left != nil {
					cur = cur.Left
				} else {
					cur = cur.Right
				}
			} else {
				if cur.Right != nil {
					cur = cur.Right
				} else {
					cur = cur.Left
				}
			}
		}
		storeValue := cur.Left.Val
		result = max(result, num^storeValue)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

