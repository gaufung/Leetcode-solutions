/*
 * @lc app=leetcode.cn id=688 lang=golang
 *
 * [688] “马”在棋盘上的概率
 */
func knightProbability(N int, K int, r int, c int) float64 {
	board := make([][][]int, K+1)
	for i := 0; i < K+1; i++ {
		board[i] = makeBoard(N)
	}
	board[0][r][c] = 1
	for k := 0; k < K; k++ {
		m := board[k]
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if m[i][j] > 0 {
					if i-2 >= 0 && j-1 >= 0 {
						board[k+1][i-2][j-1] += m[i][j]
					}
					if i-2 >= 0 && j+1 < N {
						board[k+1][i-2][j+1] += m[i][j]
					}
					if i-1 >= 0 && j-2 >= 0 {
						board[k+1][i-1][j-2] += m[i][j]
					}
					if i-1 >= 0 && j+2 < N {
						board[k+1][i-1][j+2] += m[i][j]
					}
					if i+1 < N && j-2 >= 0 {
						board[k+1][i+1][j-2] += m[i][j]
					}
					if i+1 < N && j+2 < N {
						board[k+1][i+1][j+2] += m[i][j]
					}
					if i+2 < N && j-1 >= 0 {
						board[k+1][i+2][j-1] += m[i][j]
					}
					if i+2 < N && j+1 < N {
						board[k+1][i+2][j+1] += m[i][j]
					}
				}
			}
		}
	}
	cnt := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if board[K][i][j] > 0 {
				cnt += board[K][i][j]
			}
		}
	}
	val := float64(cnt) / 1.0
	for K > 0 {
		val = val / 8.0
		K--
	}
	return val

}

func makeBoard(N int) [][]int {
	m := make([][]int, N)
	for i := 0; i < N; i++ {
		m[i] = make([]int, N)
	}
	return m
}


