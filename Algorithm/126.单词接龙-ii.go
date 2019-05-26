/*
 * @lc app=leetcode.cn id=126 lang=golang
 *
 * [126] 单词接龙 II
 *
 * https://leetcode-cn.com/problems/word-ladder-ii/description/
 *
 * algorithms
 * Hard (25.83%)
 * Likes:    52
 * Dislikes: 0
 * Total Accepted:    1.7K
 * Total Submissions: 6.5K
 * Testcase Example:  '"hit"\n"cog"\n["hot","dot","dog","lot","log","cog"]'
 *
 * 给定两个单词（beginWord 和 endWord）和一个字典 wordList，找出所有从 beginWord 到 endWord
 * 的最短转换序列。转换需遵循如下规则：
 *
 *
 * 每次转换只能改变一个字母。
 * 转换过程中的中间单词必须是字典中的单词。
 *
 *
 * 说明:
 *
 *
 * 如果不存在这样的转换序列，返回一个空列表。
 * 所有单词具有相同的长度。
 * 所有单词只由小写字母组成。
 * 字典中不存在重复的单词。
 * 你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
 *
 *
 * 示例 1:
 *
 * 输入:
 * beginWord = "hit",
 * endWord = "cog",
 * wordList = ["hot","dot","dog","lot","log","cog"]
 *
 * 输出:
 * [
 * ⁠ ["hit","hot","dot","dog","cog"],
 * ["hit","hot","lot","log","cog"]
 * ]
 *
 *
 * 示例 2:
 *
 * 输入:
 * beginWord = "hit"
 * endWord = "cog"
 * wordList = ["hot","dot","dog","lot","log"]
 *
 * 输出: []
 *
 * 解释: endWord "cog" 不在字典中，所以不存在符合要求的转换序列。
 *
 */
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	result := make([][]string, 0)
	wordList = append(wordList, beginWord)
	n := len(wordList)
	// [....., endWord, beginWord]
	neighbourhood := make(map[int][]int, 0)
	for i := 0; i < n-2; i++ {
		for j := 0; j < n-1; j++ {
			if isNeigbourhood(wordList[i], wordList[j]) {
				if val, ok := neighbourhood[i]; ok {
					neighbourhood[i] = append(val, j)
				} else {
					neighbourhood[i] = []int{j}
				}
			}
		}
	}
	for i := 0; i < n-1; i++ {
		if isNeigbourhood(wordList[n-1], wordList[i]) {
			if val, ok := neighbourhood[n-1]; ok {
				neighbourhood[n-1] = append(val, i)
			} else {
				neighbourhood[n-1] = []int{i}
			}
		}
	}
	visited := make([]bool, n)
	target := n - 2
	seed := n - 1
	stack := NewStack()
	size := n
	dfs(seed, neighbourhood, wordList, visited, &result, stack, target, &size)
	r := make([][]string, 0)
	for i := 0; i < len(result); i++ {
		if len(result[i]) == size && result[i][size-1] == endWord {
			r = append(r, result[i])
		}
	}
	return r

}

func bfs(seed int, neighbourhood map[int][]int, wordList []string, visited []int, target int) ([][]string, bool) {
	visited[seed] = true
	if seed == target {
		visited[seed] = false
		return [][]string{
			[]string{wordList[seed]},
		}, true
	}
	neighs := neighbourhood(seed)
	if len(neighs) <= 0 {
		return [][]string{}, false
	}
	result := make([][]string, 0)
	for i := 0; i < len(neighs); i++ {
		visited[neighs[i]] = true
	}

}

func createNeighbourhoold(wordList []string) map[int][]int {
	result := make(map[int][]int, 0)
	for i := 0; i < len(wordList); i++ {
		for j := i + 1; j < len(wordList); j++ {
			if isNeigbourhood(wordList[i], wordList[j]) {
				if val, ok := result[i]; ok {
					result[i] = append(val, j)
				} else {
					result[i] = []int{j}
				}
				if val, ok := result[j]; ok {
					result[j] = append(val, i)
				} else {
					result[j] = []int{i}
				}
			}
		}
	}
	return result
}

func isNeigbourhood(word1, word2 string) bool {
	diff := 0
	for i := 0; i < len(word1); i++ {
		if word1[i] != word2[i] {
			diff++
		}
	}
	return diff == 1
}

func findEndWord(wordList []string, endWord string) int {
	for i := 0; i < len(wordList); i++ {
		if wordList[i] == endWord {
			return i
		}
	}
	return -1
}

func dfs(seed int, neighbourhood map[int][]int, wordList []string, visited []bool, result *[][]string, stack *Stack, target int, size *int) {
	stack.Push(seed)
	visited[seed] = true
	if seed == target && stack.Size() <= *size {
		paths := make([]string, 0)
		for _, val := range stack.Values() {
			paths = append(paths, wordList[val])
		}
		*result = append(*result, paths)
		*size = stack.Size()
	}
	neighbour := neighbourhood[seed]
	if len(neighbour) > 0 {
		for _, neigh := range neighbour {
			if visited[neigh] == false {
				dfs(neigh, neighbourhood, wordList, visited, result, stack, target, size)
			}
		}
	}
	visited[seed] = false
	stack.Pop()
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


