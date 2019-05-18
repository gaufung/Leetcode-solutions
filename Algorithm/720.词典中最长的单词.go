import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=720 lang=golang
 *
 * [720] 词典中最长的单词
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

