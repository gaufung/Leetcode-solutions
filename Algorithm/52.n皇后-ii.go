/*
 * @lc app=leetcode.cn id=52 lang=golang
 *
 * [52] N皇后 II
 *
 * https://leetcode-cn.com/problems/n-queens-ii/description/
 *
 * algorithms
 * Hard (72.76%)
 * Likes:    52
 * Dislikes: 0
 * Total Accepted:    5.2K
 * Total Submissions: 7.1K
 * Testcase Example:  '4'
 *
 * n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
 *
 *
 *
 * 上图为 8 皇后问题的一种解法。
 *
 * 给定一个整数 n，返回 n 皇后不同的解决方案的数量。
 *
 * 示例:
 *
 * 输入: 4
 * 输出: 2
 * 解释: 4 皇后问题存在如下两个不同的解法。
 * [
 * [".Q..",  // 解法 1
 * "...Q",
 * "Q...",
 * "..Q."],
 *
 * ["..Q.",  // 解法 2
 * "Q...",
 * "...Q",
 * ".Q.."]
 * ]
 *
 *
 */
func totalNQueens(n int) int {
	result := 0
	stack := NewStack()
	stack.Push(-1)
	for stack.Size() > 0 {
		top := stack.Pop() + 1
		if top >= n {
			continue
		}
		if valid(stack, top) {
			stack.Push(top)
		} else {
			stack.Push(top)
			continue
		}
		for i := stack.Size(); i < n; i++ {
			j := 0
			for ; j < n; j++ {
				if valid(stack, j) {
					stack.Push(j)
					break
				}
			}
			if j == n {
				break
			}
		}
		if stack.Size() == n {
			result++
		}
	}
	return result
}

func valid(s *Stack, pos int) bool {
	size := s.Size()
	for i, val := range s.Values() {
		if pos == val {
			return false
		}
		if abs(size-i) == abs(val-pos) {
			return false
		}
	}
	return true
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

type Stack struct {
	elements []int
}

func NewStack() *Stack {
	return &Stack{
		elements: make([]int, 0),
	}
}

func (s *Stack) Push(val int) {
	s.elements = append(s.elements, val)
}

func (s *Stack) Pop() int {
	val := s.Top()
	s.elements = s.elements[:len(s.elements)-1]
	return val
}

func (s *Stack) Size() int {
	return len(s.elements)
}

func (s *Stack) Top() int {
	return s.elements[len(s.elements)-1]
}
func (s *Stack) Values() []int {
	return s.elements
}

