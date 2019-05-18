/*
 * @lc app=leetcode.cn id=211 lang=golang
 *
 * [211] 添加与搜索单词 - 数据结构设计
 */

type Node struct {
	Val      byte
	Terminal bool
	Next     [26]*Node
}

type WordDictionary struct {
	root *Node
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		root: &Node{},
	}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	cur := this.root
	for i := 0; i < len(word); i++ {
		letter := word[i]
		index := int(letter - 'a')
		if cur.Next[index] == nil {
			cur.Next[index] = &Node{
				Val: letter,
			}
		}
		cur = cur.Next[index]
	}
	cur.Terminal = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return this.search(this.root, word)
}

func (this *WordDictionary) search(node *Node, word string) bool {
	if len(word) == 1 {
		if word[0] == '.' {
			for i := 0; i < 26; i++ {
				if node.Next[i] != nil && node.Next[i].Terminal {
					return true
				}
			}
			return false
		} else {
			index := int(word[0] - 'a')
			if node.Next[index] != nil && node.Next[index].Terminal {
				return true
			} else {
				return false
			}
		}
	} else {
		if word[0] != '.' {
			index := int(word[0] - 'a')
			if node.Next[index] == nil {
				return false
			}
			return this.search(node.Next[index], word[1:])
		} else {
			result := false
			for i := 0; i < 26; i++ {
				if node.Next[i] != nil {
					result = result || this.search(node.Next[i], word[1:])
				}
			}
			return result
		}
	}
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

