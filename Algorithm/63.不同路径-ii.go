/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 *
 * https://leetcode-cn.com/problems/unique-paths-ii/description/
 *
 * algorithms
 * Medium (31.05%)
 * Likes:    126
 * Dislikes: 0
 * Total Accepted:    11.7K
 * Total Submissions: 37.8K
 * Testcase Example:  '[[0,0,0],[0,1,0],[0,0,0]]'
 *
 * 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
 *
 * 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
 *
 * 现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
 *
 *
 *
 * 网格中的障碍物和空位置分别用 1 和 0 来表示。
 *
 * 说明：m 和 n 的值均不超过 100。
 *
 * 示例 1:
 *
 * 输入:
 * [
 * [0,0,0],
 * [0,1,0],
 * [0,0,0]
 * ]
 * 输出: 2
 * 解释:
 * 3x3 网格的正中间有一个障碍物。
 * 从左上角到右下角一共有 2 条不同的路径：
 * 1. 向右 -> 向右 -> 向下 -> 向下
 * 2. 向下 -> 向下 -> 向右 -> 向右
 *
 *
 */
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m <= 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n <= 0 {
		return 0
	}
	paths := make([][]int, m)
	for i := 0; i < m; i++ {
		paths[i] = make([]int, n)
	}
	if obstacleGrid[0][0] == 1 {
		paths[0][0] = 0
	} else {
		paths[0][0] = 1
	}
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			paths[0][j] = 0
		} else {
			paths[0][j] = paths[0][j-1]
		}
	}
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			paths[i][0] = 0
		} else {
			paths[i][0] = paths[i-1][0]
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				paths[i][j] = 0
			} else {
				paths[i][j] = paths[i-1][j] + paths[i][j-1]
			}
		}
	}
	return paths[m-1][n-1]
}

