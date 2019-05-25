/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] 单词搜索 II
 *
 * https://leetcode-cn.com/problems/word-search-ii/description/
 *
 * algorithms
 * Hard (33.37%)
 * Likes:    35
 * Dislikes: 0
 * Total Accepted:    2.5K
 * Total Submissions: 7.5K
 * Testcase Example:  '[["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]]\n["oath","pea","eat","rain"]'
 *
 * 给定一个二维网格 board 和一个字典中的单词列表 words，找出所有同时在二维网格和字典中出现的单词。
 *
 *
 * 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。
 *
 * 示例:
 *
 * 输入:
 * words = ["oath","pea","eat","rain"] and board =
 * [
 * ⁠ ['o','a','a','n'],
 * ⁠ ['e','t','a','e'],
 * ⁠ ['i','h','k','r'],
 * ⁠ ['i','f','l','v']
 * ]
 *
 * 输出: ["eat","oath"]
 *
 * 说明:
 * 你可以假设所有输入都由小写字母 a-z 组成。
 *
 * 提示:
 *
 *
 * 你需要优化回溯算法以通过更大数据量的测试。你能否早点停止回溯？
 * 如果当前单词不存在于所有单词的前缀中，则可以立即停止回溯。什么样的数据结构可以有效地执行这样的操作？散列表是否可行？为什么？
 * 前缀树如何？如果你想学习如何实现一个基本的前缀树，请先查看这个问题： 实现Trie（前缀树）。
 *
 *
 */
func findWords(board [][]byte, words []string) []string {
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	root := NewTrieNode('0')
	for _, word := range words {
		InsertWord(root, word)
	}
	bytes := make([]byte, 0)
	result := make(map[string]bool, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if root.Next[int(board[i][j]-'a')] != nil {
				dfs(i, j, m, n, board, root.Next[int(board[i][j]-'a')], visited, &bytes, result)
			}

		}
	}
	res := make([]string, 0)
	for k, _ := range result {
		res = append(res, k)
	}
	return res
}

func dfs(i, j int, m, n int, board [][]byte, node *TrieNode, visited [][]bool, bytes *[]byte, result map[string]bool) {
	visited[i][j] = true
	*bytes = append(*bytes, board[i][j])
	if node.Teriminal {
		result[string(*bytes)] = true
	}
	if i-1 >= 0 && visited[i-1][j] == false {
		letter := board[i-1][j]
		index := int(letter - 'a')
		if node.Next[index] != nil {
			dfs(i-1, j, m, n, board, node.Next[index], visited, bytes, result)
		}
	}
	if i+1 < m && visited[i+1][j] == false {
		letter := board[i+1][j]
		index := int(letter - 'a')
		if node.Next[index] != nil {
			dfs(i+1, j, m, n, board, node.Next[index], visited, bytes, result)
		}
	}
	if j-1 >= 0 && visited[i][j-1] == false {
		letter := board[i][j-1]
		index := int(letter - 'a')
		if node.Next[index] != nil {
			dfs(i, j-1, m, n, board, node.Next[index], visited, bytes, result)
		}
	}
	if j+1 < n && visited[i][j+1] == false {
		letter := board[i][j+1]
		index := int(letter - 'a')
		if node.Next[index] != nil {
			dfs(i, j+1, m, n, board, node.Next[index], visited, bytes, result)
		}
	}
	visited[i][j] = false
	*bytes = (*bytes)[:len(*bytes)-1]
}

type TrieNode struct {
	Letter    byte
	str       string
	Teriminal bool
	Next      [26]*TrieNode
}

func NewTrieNode(letter byte) *TrieNode {
	return &TrieNode{
		Letter: letter,
		str:    string(letter),
	}
}

func InsertWord(root *TrieNode, word string) {
	cur := root
	for i := 0; i < len(word); i++ {
		index := int(word[i] - 'a')
		if cur.Next[index] == nil {
			cur.Next[index] = NewTrieNode(word[i])
		}
		cur = cur.Next[index]
	}
	cur.Teriminal = true
}

