/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 *
 * https://leetcode-cn.com/problems/word-ladder/description/
 *
 * algorithms
 * Medium (33.97%)
 * Likes:    85
 * Dislikes: 0
 * Total Accepted:    5.7K
 * Total Submissions: 16.8K
 * Testcase Example:  '"hit"\n"cog"\n["hot","dot","dog","lot","log","cog"]'
 *
 * 给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord
 * 的最短转换序列的长度。转换需遵循如下规则：
 *
 *
 * 每次转换只能改变一个字母。
 * 转换过程中的中间单词必须是字典中的单词。
 *
 *
 * 说明:
 *
 *
 * 如果不存在这样的转换序列，返回 0。
 * 所有单词具有相同的长度。
 * 所有单词只由小写字母组成。
 * 字典中不存在重复的单词。
 * 你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
 *
 *
 * 示例 1:
 *
 * 输入:
 * beginWord = "hit",
 * endWord = "cog",
 * wordList = ["hot","dot","dog","lot","log","cog"]
 *
 * 输出: 5
 *
 * 解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
 * ⁠    返回它的长度 5。
 *
 *
 * 示例 2:
 *
 * 输入:
 * beginWord = "hit"
 * endWord = "cog"
 * wordList = ["hot","dot","dog","lot","log"]
 *
 * 输出: 0
 *
 * 解释: endWord "cog" 不在字典中，所以无法进行转换。
 *
 */
func ladderLength(beginWord string, endWord string, wordList []string) int {
	endWordIndex := findIndex(wordList, endWord)
	if endWordIndex == -1 {
		return 0
	}
	wordList = append(wordList, beginWord)
	n := len(wordList)
	connections := make([][]int, n)
	for i := 0; i < n; i++ {
		connections[i] = make([]int, n)
		for j := i; j < n; j++ {
			if wordsAdjacent(wordList[i], wordList[j]) {
				connections[i][j] = 1
			}
		}
	}
	steps := make([]int, n)
	visited := make([]bool, n)
	steps[n-1] = 0
	visited[n-1] = true
	queue := make([]int, 1)
	queue[0] = n - 1
	currentStep := 1
	for len(queue) > 0 {
		currentStep++
		newQueue := make([]int, 0)
		for _, q := range queue {
			for i := 0; i < n; i++ {
				if (connections[i][q] == 1 || connections[q][i] == 1) && visited[i] == false {
					steps[i] = currentStep
					visited[i] = true
					newQueue = append(newQueue, i)
				}
			}
		}
		queue = newQueue
	}
	return steps[endWordIndex]
}

func findIndex(wordList []string, word string) int {
	for i := 0; i < len(wordList); i++ {
		if wordList[i] == word {
			return i
		}
	}
	return -1
}

func wordsAdjacent(word1, word2 string) bool {
	cnt := 0
	i, n := 0, len(word1)
	for i < n {
		if word1[i] != word2[i] {
			cnt++
		}
		i++
	}
	return cnt == 1
}

