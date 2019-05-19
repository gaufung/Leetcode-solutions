/*
 * @lc app=leetcode.cn id=677 lang=golang
 *
 * [677] 键值映射
 *
 * https://leetcode-cn.com/problems/map-sum-pairs/description/
 *
 * algorithms
 * Medium (58.28%)
 * Likes:    23
 * Dislikes: 0
 * Total Accepted:    2.1K
 * Total Submissions: 3.6K
 * Testcase Example:  '["MapSum", "insert", "sum", "insert", "sum"]\n[[], ["apple",3], ["ap"], ["app",2], ["ap"]]'
 *
 * 实现一个 MapSum 类里的两个方法，insert 和 sum。
 *
 * 对于方法 insert，你将得到一对（字符串，整数）的键值对。字符串表示键，整数表示值。如果键已经存在，那么原来的键值对将被替代成新的键值对。
 *
 * 对于方法 sum，你将得到一个表示前缀的字符串，你需要返回所有以该前缀开头的键的值的总和。
 *
 * 示例 1:
 *
 * 输入: insert("apple", 3), 输出: Null
 * 输入: sum("ap"), 输出: 3
 * 输入: insert("app", 2), 输出: Null
 * 输入: sum("ap"), 输出: 5
 *
 *
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

