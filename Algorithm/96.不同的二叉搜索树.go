/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 *
 * https://leetcode-cn.com/problems/unique-binary-search-trees/description/
 *
 * algorithms
 * Medium (60.04%)
 * Likes:    147
 * Dislikes: 0
 * Total Accepted:    8.2K
 * Total Submissions: 13.6K
 * Testcase Example:  '3'
 *
 * 给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？
 *
 * 示例:
 *
 * 输入: 3
 * 输出: 5
 * 解释:
 * 给定 n = 3, 一共有 5 种不同结构的二叉搜索树:
 *
 * ⁠  1         3     3      2      1
 * ⁠   \       /     /      / \      \
 * ⁠    3     2     1      1   3      2
 * ⁠   /     /       \                 \
 * ⁠  2     1         2                 3
 *
 */
func numTrees(n int) int {
	if n <= 0 {
		return 0
	}
	return count(1, n)
}
func count(from, to int) int {
	if from >= to {
		return 1
	}
	sum := 0
	for i := from; i <= to; i++ {
		leftCount, rightCount := 1, 1
		if i != from {
			leftCount = count(from, i-1)
		}
		if i != to {
			rightCount = count(i+1, to)
		}
		sum += leftCount * rightCount
	}
	return sum
}

