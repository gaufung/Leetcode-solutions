import "sort"

/*
 * @lc app=leetcode.cn id=332 lang=golang
 *
 * [332] 重新安排行程
 *
 * https://leetcode-cn.com/problems/reconstruct-itinerary/description/
 *
 * algorithms
 * Medium (30.21%)
 * Likes:    33
 * Dislikes: 0
 * Total Accepted:    959
 * Total Submissions: 3.2K
 * Testcase Example:  '[["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]'
 *
 * 给定一个机票的字符串二维数组 [from,
 * to]，子数组中的两个成员分别表示飞机出发和降落的机场地点，对该行程进行重新规划排序。所有这些机票都属于一个从JFK（肯尼迪国际机场）出发的先生，所以该行程必须从
 * JFK 出发。
 *
 * 说明:
 *
 *
 * 如果存在多种有效的行程，你可以按字符自然排序返回最小的行程组合。例如，行程 ["JFK", "LGA"] 与 ["JFK", "LGB"]
 * 相比就更小，排序更靠前
 * 所有的机场都用三个大写字母表示（机场代码）。
 * 假定所有机票至少存在一种合理的行程。
 *
 *
 * 示例 1:
 *
 * 输入: [["MUC", "LHR"], ["JFK", "MUC"], ["SFO", "SJC"], ["LHR", "SFO"]]
 * 输出: ["JFK", "MUC", "LHR", "SFO", "SJC"]
 *
 *
 * 示例 2:
 *
 * 输入: [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
 * 输出: ["JFK","ATL","JFK","SFO","ATL","SFO"]
 * 解释: 另一种有效的行程是 ["JFK","SFO","ATL","JFK","ATL","SFO"]。但是它自然排序更大更靠后。
 *
 */
func findItinerary(tickets [][]string) []string {
	sort.Slice(tickets, func(i, j int) bool {
		if tickets[i][0] < tickets[j][0] {
			return true
		} else if tickets[i][0] > tickets[j][0] {
			return false
		} else {
			return tickets[i][1] < tickets[j][1]
		}
	})
	n := len(tickets)
	if n <= 0 {
		return []string{}
	}
	s := NewStack()
	vistied := make([]bool, n)
	s.Push("JFK")
	dfs("JFK", tickets, vistied, s)
	return s.Value()
}

func dfs(start string, tickets [][]string, vistied []bool, stack *Stack) {
	for i := 0; i < len(tickets); i++ {
		if vistied[i] == false && tickets[i][0] == start {
			vistied[i] = true
			stack.Push(tickets[i][1])
			dfs(tickets[i][1], tickets, vistied, stack)
			if stack.Size() != len(tickets)+1 {
				vistied[i] = false
				stack.Pop()
			}
		}
	}
}

type Stack struct {
	elements []string
}

func NewStack() *Stack {
	return &Stack{
		elements: make([]string, 0),
	}
}

func (s *Stack) Size() int {
	return len(s.elements)
}

func (s *Stack) Push(node string) {
	s.elements = append(s.elements, node)
}

func (s *Stack) Pop() string {
	size := s.Size()
	val := s.elements[size-1]
	s.elements = s.elements[:size-1]
	return val
}

func (s *Stack) Top() string {
	size := s.Size()
	return s.elements[size-1]
}

func (s *Stack) Value() []string {
	return s.elements
}
