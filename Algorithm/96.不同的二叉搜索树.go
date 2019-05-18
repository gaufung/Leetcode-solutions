/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
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

