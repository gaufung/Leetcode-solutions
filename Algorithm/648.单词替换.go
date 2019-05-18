import "strings"

/*
 * @lc app=leetcode.cn id=648 lang=golang
 *
 * [648] 单词替换
 */
func replaceWords(dict []string, sentence string) string {
	t := NewTrieTree()
	for _, word := range dict {
		t.InsertWord(word)
	}
	words := strings.Split(sentence, " ")
	filterWords := make([]string, len(words))
	for i, word := range words {
		filterWords[i] = t.FindRoot(word)
	}
	return strings.Join(filterWords, " ")
}

type TrieNode struct {
	Letter   byte
	Terminal bool
	Next     [26]*TrieNode
}

type TrieTree struct {
	root *TrieNode
}

func NewTrieTree() *TrieTree {
	return &TrieTree{
		root: &TrieNode{},
	}
}

func (t *TrieTree) InsertWord(word string) {
	cur := t.root
	for i := 0; i < len(word); i++ {
		index := int(word[i] - 'a')
		if cur.Next[index] == nil {
			cur.Next[index] = &TrieNode{Letter: word[i]}
		}
		cur = cur.Next[index]
	}
	cur.Terminal = true
}

func (t *TrieTree) FindRoot(word string) string {
	cur := t.root
	i := 0
	for i < len(word) && cur != nil && cur.Terminal == false {
		index := int(word[i] - 'a')
		cur = cur.Next[index]
		i++
	}
	if cur == nil {
		return word
	}
	if cur.Terminal == true {
		return word[:i]
	} else {
		return word
	}

}

