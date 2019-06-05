import "strconv"

/*
 * @lc app=leetcode.cn id=655 lang=golang
 *
 * [655] 输出二叉树
 *
 * https://leetcode-cn.com/problems/print-binary-tree/description/
 *
 * algorithms
 * Medium (48.11%)
 * Likes:    24
 * Dislikes: 0
 * Total Accepted:    904
 * Total Submissions: 1.9K
 * Testcase Example:  '[1,2]'
 *
 * 在一个 m*n 的二维字符串数组中输出二叉树，并遵守以下规则：
 *
 *
 * 行数 m 应当等于给定二叉树的高度。
 * 列数 n 应当总是奇数。
 *
 * 根节点的值（以字符串格式给出）应当放在可放置的第一行正中间。根节点所在的行与列会将剩余空间划分为两部分（左下部分和右下部分）。你应该将左子树输出在左下部分，右子树输出在右下部分。左下和右下部分应当有相同的大小。即使一个子树为空而另一个非空，你不需要为空的子树输出任何东西，但仍需要为另一个子树留出足够的空间。然而，如果两个子树都为空则不需要为它们留出任何空间。
 * 每个未使用的空间应包含一个空的字符串""。
 * 使用相同的规则输出子树。
 *
 *
 * 示例 1:
 *
 *
 * 输入:
 * ⁠    1
 * ⁠   /
 * ⁠  2
 * 输出:
 * [["", "1", ""],
 * ⁠["2", "", ""]]
 *
 *
 * 示例 2:
 *
 *
 * 输入:
 * ⁠    1
 * ⁠   / \
 * ⁠  2   3
 * ⁠   \
 * ⁠    4
 * 输出:
 * [["", "", "", "1", "", "", ""],
 * ⁠["", "2", "", "", "", "3", ""],
 * ⁠["", "", "4", "", "", "", ""]]
 *
 *
 * 示例 3:
 *
 *
 * 输入:
 * ⁠     1
 * ⁠    / \
 * ⁠   2   5
 * ⁠  /
 * ⁠ 3
 * ⁠/
 * 4
 * 输出:
 * [["",  "",  "", "",  "", "", "", "1", "",  "",  "",  "",  "", "", ""]
 * ⁠["",  "",  "", "2", "", "", "", "",  "",  "",  "",  "5", "", "", ""]
 * ⁠["",  "3", "", "",  "", "", "", "",  "",  "",  "",  "",  "", "", ""]
 * ⁠["4", "",  "", "",  "", "", "", "",  "",  "",  "",  "",  "", "", ""]]
 *
 *
 * 注意: 二叉树的高度在范围 [1, 10] 中。
 *
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func printTree(root *TreeNode) [][]string {
	if root == nil {
		return [][]string{}
	}
	hi := height(root)
	result := make([][]string, hi)
	for i := 0; i < hi; i++ {
		result[i] = make([]string, 1<<(uint(hi))-1)
	}
	print(root, 0, 0, 1<<(uint(hi))-2, result)
	return result

}

func print(node *TreeNode, level, lo, hi int, result [][]string) {
	if node == nil {
		return
	}
	mi := (lo + hi) / 2
	result[level][mi] = strconv.Itoa(node.Val)
	print(node.Left, level+1, lo, mi-1, result)
	print(node.Right, level+1, mi+1, hi, result)
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	} else {
		return 1 + max(height(node.Left), height(node.Right))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

