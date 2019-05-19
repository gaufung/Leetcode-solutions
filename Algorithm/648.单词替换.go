import "strings"

/*
 * @lc app=leetcode.cn id=648 lang=golang
 *
 * [648] 单词替换
 *
 * https://leetcode-cn.com/problems/replace-words/description/
 *
 * algorithms
 * Medium (50.14%)
 * Likes:    27
 * Dislikes: 0
 * Total Accepted:    1.5K
 * Total Submissions: 2.9K
 * Testcase Example:  '["cat", "bat", "rat"]\n"the cattle was rattled by the battery"'
 *
 * 在英语中，我们有一个叫做 词根(root)的概念，它可以跟着其他一些词组成另一个较长的单词——我们称这个词为
 * 继承词(successor)。例如，词根an，跟随着单词 other(其他)，可以形成新的单词 another(另一个)。
 *
 * 现在，给定一个由许多词根组成的词典和一个句子。你需要将句子中的所有继承词用词根替换掉。如果继承词有许多可以形成它的词根，则用最短的词根替换它。
 *
 * 你需要输出替换之后的句子。
 *
 * 示例 1:
 *
 *
 * 输入: dict(词典) = ["cat", "bat", "rat"]
 * sentence(句子) = "the cattle was rattled by the battery"
 * 输出: "the cat was rat by the bat"
 *
 *
 * 注:
 *
 *
 * 输入只包含小写字母。
 * 1 <= 字典单词数 <=1000
 * 1 <=  句中词语数 <= 1000
 * 1 <= 词根长度 <= 100
 * 1 <= 句中词语长度 <= 1000
 *
 *
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

