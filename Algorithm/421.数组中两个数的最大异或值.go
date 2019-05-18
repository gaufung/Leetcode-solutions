/*
 * @lc app=leetcode.cn id=421 lang=golang
 *
 * [421] 数组中两个数的最大异或值
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
