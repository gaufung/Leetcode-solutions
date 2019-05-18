/*
 * @lc app=leetcode.cn id=677 lang=golang
 *
 * [677] 键值映射
 */
type MapSum struct {
	root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{
		root: &TrieNode{},
	}
}

func (this *MapSum) Insert(key string, val int) {
	cur := this.root
	for i := 0; i < len(key); i++ {
		index := int(key[i] - 'a')
		if cur.next[index] == nil {
			cur.next[index] = &TrieNode{
				letter: key[i],
			}
		}
		cur = cur.next[index]
	}
	cur.value = val
	cur.terminal = true
}

func (this *MapSum) Sum(prefix string) int {
	cur := this.root
	for i := 0; i < len(prefix); i++ {
		index := int(prefix[i] - 'a')
		if cur.next[index] == nil {
			cur = cur.next[index]
			break
		} else {
			cur = cur.next[index]
		}
	}
	return this.sum(cur)
}

func (this *MapSum) sum(node *TrieNode) int {
	if node == nil {
		return 0
	}
	result := node.value
	for i := 0; i < 26; i++ {
		result += this.sum(node.next[i])
	}
	return result
}

type TrieNode struct {
	letter   byte
	terminal bool
	value    int
	next     [26]*TrieNode
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */

