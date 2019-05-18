/*
 * @lc app=leetcode.cn id=315 lang=golang
 *
 * [315] 计算右侧小于当前元素的个数
 */
func countSmaller(nums []int) []int {
	n := len(nums)
	if n <= 0 {
		return []int{}
	}
	result := make([]int, n)
	root := &TreeNode{Val: nums[n-1]}
	result[n-1] = 0
	for i := n - 2; i >= 0; i-- {
		result[i] = Insert(root, &TreeNode{Val: nums[i]})
	}
	return result
}

func Insert(root *TreeNode, node *TreeNode) int {
	for {
		if node.Val > root.Val && root.Right != nil {
			node.Less = node.Less + 1
			root = root.Right
			continue
		}
		if node.Val > root.Val && root.Right == nil {
			node.Less = node.Less + 1
			root.Right = node
			break
		}
		if node.Val < root.Val && root.Left != nil {
			root = root.Left
			continue
		}
		if node.Val < root.Val && root.Left == nil {
			root.Less = node
			break
		}
	}
	return node.Less
}

type TreeNode struct {
	Val   int
	Less  int
	Left  *TreeNode
	Right *TreeNode
}

