import "sort"

/*
 * @lc app=leetcode.cn id=720 lang=golang
 *
 * [720] 词典中最长的单词
 *
 * https://leetcode-cn.com/problems/longest-word-in-dictionary/description/
 *
 * algorithms
 * Easy (42.04%)
 * Likes:    33
 * Dislikes: 0
 * Total Accepted:    2.3K
 * Total Submissions: 5.5K
 * Testcase Example:  '["w","wo","wor","worl","world"]'
 *
 *
 * 给出一个字符串数组words组成的一本英语词典。从中找出最长的一个单词，该单词是由words词典中其他单词逐步添加一个字母组成。若其中有多个可行的答案，则返回答案中字典序最小的单词。
 *
 * 若无答案，则返回空字符串。
 *
 * 示例 1:
 *
 *
 * 输入:
 * words = ["w","wo","wor","worl", "world"]
 * 输出: "world"
 * 解释:
 * 单词"world"可由"w", "wo", "wor", 和 "worl"添加一个字母组成。
 *
 *
 * 示例 2:
 *
 *
 * 输入:
 * words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
 * 输出: "apple"
 * 解释:
 * "apply"和"apple"都能由词典中的单词组成。但是"apple"得字典序小于"apply"。
 *
 *
 * 注意:
 *
 *
 * 所有输入的字符串都只包含小写字母。
 * words数组长度范围为[1,1000]。
 * words[i]的长度范围为[1,30]。
 *
 *
 */
func longestWord(words []string) string {
	sort.Slice(words, func(i, j int) bool { return len(words[i]) < len(words[j]) })
	root := &TrieNode{
		terminal: true,
	}
	Insert(root, words[0])
	candidate := words[0]
	for i := 1; i < len(words); i++ {
		length := Insert(root, words[i])
		if length > len(candidate) {
			candidate = words[i]
		}
		if length == len(candidate) {
			if words[i] < candidate {
				candidate = words[i]
			}
		}
	}
	return candidate
}

type TrieNode struct {
	letter   byte
	terminal bool
	next     [26]*TrieNode
}

func Insert(root *TrieNode, word string) int {
	cur := root
	isNext := true
	for i := 0; i < len(word); i++ {
		index := int(word[i] - 'a')
		if cur.next[index] == nil {
			cur.next[index] = &TrieNode{letter: word[i]}
		}
		if cur.terminal == false {
			isNext = false
		}
		cur = cur.next[index]
	}
	cur.terminal = true
	if isNext {
		return len(word)
	} else {
		return -1
	}
}

