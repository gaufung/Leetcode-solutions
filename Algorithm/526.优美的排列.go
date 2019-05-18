/*
 * @lc app=leetcode.cn id=526 lang=golang
 *
 * [526] 优美的排列
 */
func countArrangement(N int) int {
	nums := make([]int, N)
	for i := 0; i < N; i++ {
		nums[i] = i + 1
	}
	return arrange(nums, 1)
}

func arrange(nums []int, startIndex int) int {
	if len(nums) == 1 {
		if nums[0]%startIndex == 0 {
			return 1
		}
		if startIndex%nums[0] == 0 {
			return 1
		}
		return 0
	}
	result := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%startIndex == 0 || startIndex%nums[i] == 0 {
			leftNums := removeKth(nums, i)
			result += arrange(leftNums, startIndex+1)
		}
	}
	return result
}

func removeKth(nums []int, k int) []int {
	result := make([]int, 0)
	for i, val := range nums {
		if i != k {
			result = append(result, val)
		}
	}
	return result
}

// type Stack struct {
// 	elements []int
// }

// func NewStack() *Stack {
// 	return &Stack{
// 		elements: make([]int, 0),
// 	}
// }

// func (s *Stack) Push(val int) {
// 	s.elements = append(s.elements, val)
// }

// func (s *Stack) Pop() int {
// 	val := s.Top()
// 	s.elements = s.elements[:len(s.elements)-1]
// 	return val
// }

// func (s *Stack) Size() int {
// 	return len(s.elements)
// }

// func (s *Stack) Top() int {
// 	return s.elements[len(s.elements)-1]
// }
// func (s *Stack) Values() []int {
// 	return s.elements
// }