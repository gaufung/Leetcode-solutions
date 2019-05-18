/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
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

