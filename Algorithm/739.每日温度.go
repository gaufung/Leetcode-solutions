/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 *
 * https://leetcode-cn.com/problems/daily-temperatures/description/
 *
 * algorithms
 * Medium (57.74%)
 * Likes:    99
 * Dislikes: 0
 * Total Accepted:    5.8K
 * Total Submissions: 10.1K
 * Testcase Example:  '[73,74,75,71,69,72,76,73]'
 *
 * 根据每日 气温 列表，请重新生成一个列表，对应位置的输入是你需要再等待多久温度才会升高的天数。如果之后都不会升高，请输入 0 来代替。
 *
 * 例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4,
 * 2, 1, 1, 0, 0]。
 *
 * 提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的都是 [30, 100] 范围内的整数。
 *
 */
func dailyTemperatures(T []int) []int {
	n := len(T)
	if n <= 0 {
		return []int{}
	}
	result := make([]int, n)
	stack := NewStack()
	stack.Push(0)
	for i := 1; i < n; i++ {
		if stack.Size() == 0 {
			stack.Push(i)
		} else {
			if T[i] <= T[stack.Top()] {
				stack.Push(i)
			} else {
				for stack.Size() > 0 && T[i] > T[stack.Top()] {
					index := stack.Pop()
					result[index] = i - index
				}
				stack.Push(i)
			}
		}
	}
	return result
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

func (s *Stack) Size() int {
	return len(s.elements)
}

func (s *Stack) Pop() int {
	val := s.elements[s.Size()-1]
	s.elements = s.elements[:s.Size()-1]
	return val
}

func (s *Stack) Top() int {
	return s.elements[s.Size()-1]
}

