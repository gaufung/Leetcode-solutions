/*
 * @lc app=leetcode.cn id=390 lang=golang
 *
 * [390] 消除游戏
 *
 * https://leetcode-cn.com/problems/elimination-game/description/
 *
 * algorithms
 * Medium (37.06%)
 * Likes:    28
 * Dislikes: 0
 * Total Accepted:    868
 * Total Submissions: 2.3K
 * Testcase Example:  '9'
 *
 * 给定一个从1 到 n 排序的整数列表。
 * 首先，从左到右，从第一个数字开始，每隔一个数字进行删除，直到列表的末尾。
 * 第二步，在剩下的数字中，从右到左，从倒数第一个数字开始，每隔一个数字进行删除，直到列表开头。
 * 我们不断重复这两步，从左到右和从右到左交替进行，直到只剩下一个数字。
 * 返回长度为 n 的列表中，最后剩下的数字。
 *
 * 示例：
 *
 *
 * 输入:
 * n = 9,
 * 1 2 3 4 5 6 7 8 9
 * 2 4 6 8
 * 2 6
 * 6
 *
 * 输出:
 * 6
 *
 */
func lastRemaining(n int) int {
	l := NewList()
	for i := 1; i <= n; i++ {
		l.Push(i)
	}
	isFromHead := true
	for l.Size() > 1 {
		if isFromHead {
			cur := l.head.Next
			for cur != nil && !l.IsTail(cur) {
				l.Remove(cur)
				cur = cur.Next.Next
			}
		} else {
			cur := l.tail.Prev
			for cur != nil && !l.IsHead(cur) {
				l.Remove(cur)
				cur = cur.Prev.Prev
			}
		}
		isFromHead = !isFromHead
	}
	return l.head.Next.Val
}

type Element struct {
	Val  int
	Prev *Element
	Next *Element
}

type List struct {
	head *Element
	tail *Element
	size int
}

func NewList() *List {
	l := new(List)
	l.head = &Element{}
	l.tail = &Element{}
	l.head.Next = l.tail
	l.tail.Prev = l.head
	return l
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Remove(e *Element) {
	e.Prev.Next = e.Next
	e.Next.Prev = e.Prev
	l.size--
}

func (l *List) Push(val int) {
	e := &Element{Val: val}
	l.tail.Prev.Next = e
	e.Prev = l.tail.Prev
	e.Next = l.tail
	l.tail.Prev = e
	l.size++
}

func (l *List) IsHead(e *Element) bool {
	return e == l.head
}

func (l *List) IsTail(e *Element) bool {
	return e == l.tail
}



