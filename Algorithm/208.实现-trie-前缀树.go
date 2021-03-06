/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 *
 * https://leetcode-cn.com/problems/implement-trie-prefix-tree/description/
 *
 * algorithms
 * Medium (57.93%)
 * Likes:    67
 * Dislikes: 0
 * Total Accepted:    6.5K
 * Total Submissions: 11.2K
 * Testcase Example:  '["Trie","insert","search","search","startsWith","insert","search"]\n[[],["apple"],["apple"],["app"],["app"],["app"],["app"]]'
 *
 * 实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。
 *
 * 示例:
 *
 * Trie trie = new Trie();
 *
 * trie.insert("apple");
 * trie.search("apple");   // 返回 true
 * trie.search("app");     // 返回 false
 * trie.startsWith("app"); // 返回 true
 * trie.insert("app");
 * trie.search("app");     // 返回 true
 *
 * 说明:
 *
 *
 * 你可以假设所有的输入都是由小写字母 a-z 构成的。
 * 保证所有输入均为非空字符串。
 *
 *
 */
type Trie struct {
	letter      byte
	termination bool
	next        [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	t := this
	for _, letter := range word {
		index := int(letter - 'a')
		if t.next[index] == nil {
			tt := Constructor()
			tt.letter = byte(letter)
			t.next[index] = &tt
		}
		t = t.next[index]
	}
	t.termination = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	t := this
	for _, letter := range word {
		index := int(letter - 'a')
		if t.next[index] != nil {
			t = t.next[index]
		} else {
			return false
		}
	}
	return t.termination
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	t := this
	for _, letter := range prefix {
		index := int(letter - 'a')
		if t.next[index] != nil {
			t = t.next[index]
		} else {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

