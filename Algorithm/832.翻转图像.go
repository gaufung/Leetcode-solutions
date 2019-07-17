/*
 * @lc app=leetcode.cn id=832 lang=golang
 *
 * [832] 翻转图像
 *
 * https://leetcode-cn.com/problems/flipping-an-image/description/
 *
 * algorithms
 * Easy (71.96%)
 * Likes:    118
 * Dislikes: 0
 * Total Accepted:    17.5K
 * Total Submissions: 24.2K
 * Testcase Example:  '[[1,1,0],[1,0,1],[0,0,0]]'
 *
 * 给定一个二进制矩阵 A，我们想先水平翻转图像，然后反转图像并返回结果。
 *
 * 水平翻转图片就是将图片的每一行都进行翻转，即逆序。例如，水平翻转 [1, 1, 0] 的结果是 [0, 1, 1]。
 *
 * 反转图片的意思是图片中的 0 全部被 1 替换， 1 全部被 0 替换。例如，反转 [0, 1, 1] 的结果是 [1, 0, 0]。
 *
 * 示例 1:
 *
 *
 * 输入: [[1,1,0],[1,0,1],[0,0,0]]
 * 输出: [[1,0,0],[0,1,0],[1,1,1]]
 * 解释: 首先翻转每一行: [[0,1,1],[1,0,1],[0,0,0]]；
 * ⁠    然后反转图片: [[1,0,0],[0,1,0],[1,1,1]]
 *
 *
 * 示例 2:
 *
 *
 * 输入: [[1,1,0,0],[1,0,0,1],[0,1,1,1],[1,0,1,0]]
 * 输出: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
 * 解释: 首先翻转每一行: [[0,0,1,1],[1,0,0,1],[1,1,1,0],[0,1,0,1]]；
 * ⁠    然后反转图片: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
 *
 *
 * 说明:
 *
 *
 * 1 <= A.length = A[0].length <= 20
 * 0 <= A[i][j] <= 1
 *
 *
 */
func flipAndInvertImage(A [][]int) [][]int {
	M := len(A)
	if M <= 0 {
		return A
	}
	N := len(A[0])
	if N <= 0 {
		return A
	}
	for i := 0; i < M; i++ {
		for j, k := 0, N-1; j <= k; {
			if j == k {
				A[i][j] = inverse(A[i][j])
			} else {
				A[i][j] = inverse(A[i][j])
				A[i][k] = inverse(A[i][k])
				A[i][j], A[i][k] = A[i][k], A[i][j]
			}
			j++
			k--
		}
	}
	return A
}

func inverse(a int) int {
	if a == 0 {
		return 1
	}
	return 0
}

