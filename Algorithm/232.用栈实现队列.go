/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 *
 * https://leetcode-cn.com/problems/implement-queue-using-stacks/description/
 *
 * algorithms
 * Easy (59.95%)
 * Likes:    72
 * Dislikes: 0
 * Total Accepted:    11.8K
 * Total Submissions: 19.7K
 * Testcase Example:  '["MyQueue","push","push","peek","pop","empty"]\n[[],[1],[2],[],[],[]]'
 *
 * 使用栈实现队列的下列操作：
 *
 *
 * push(x) -- 将一个元素放入队列的尾部。
 * pop() -- 从队列首部移除元素。
 * peek() -- 返回队列首部的元素。
 * empty() -- 返回队列是否为空。
 *
 *
 * 示例:
 *
 * MyQueue queue = new MyQueue();
 *
 * queue.push(1);
 * queue.push(2);
 * queue.peek();  // 返回 1
 * queue.pop();   // 返回 1
 * queue.empty(); // 返回 false
 *
 * 说明:
 *
 *
 * 你只能使用标准的栈操作 -- 也就是只有 push to top, peek/pop from top, size, 和 is empty
 * 操作是合法的。
 * 你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
 * 假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）。
 *
 *
 */
type MyQueue struct {
	wareHouse MyStack
	backup    MyStack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		wareHouse: NewMyStack(),
		backup:    NewMyStack(),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.wareHouse.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	for this.wareHouse.Size() > 1 {
		this.backup.Push(this.wareHouse.Pop())
	}
	val := this.wareHouse.Pop()
	this.backup, this.wareHouse = this.wareHouse, this.backup
	return val
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	for this.wareHouse.Size() > 1 {
		this.backup.Push(this.wareHouse.Pop())
	}
	val := this.wareHouse.Pop()
	this.backup.Push(val)
	this.backup, this.wareHouse = this.wareHouse, this.backup
	return val
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.wareHouse.Size() == 0
}

type MyStack struct {
	elements []int
}

func NewMyStack() MyStack {
	return MyStack{
		elements: make([]int, 0),
	}
}

func (this *MyStack) Push(val int) {
	this.elements = append(this.elements, val)
}

func (this *MyStack) Peek() int {
	n := len(this.elements)
	return this.elements[n-1]
}

func (this *MyStack) Pop() int {
	val := this.Peek()
	n := len(this.elements)
	this.elements = this.elements[0 : n-1]
	return val
}

func (this *MyStack) Size() int {
	return len(this.elements)
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

