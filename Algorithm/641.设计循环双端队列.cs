/*
 * @lc app=leetcode.cn id=641 lang=csharp
 *
 * [641] 设计循环双端队列
 *
 * https://leetcode-cn.com/problems/design-circular-deque/description/
 *
 * algorithms
 * Medium (45.85%)
 * Likes:    15
 * Dislikes: 0
 * Total Accepted:    1.6K
 * Total Submissions: 3.5K
 * Testcase Example:  '["MyCircularDeque","insertLast","insertLast","insertFront","insertFront","getRear","isFull","deleteLast","insertFront","getFront"]\n[[3],[1],[2],[3],[4],[],[],[],[4],[]]'
 *
 * 设计实现双端队列。
 * 你的实现需要支持以下操作：
 * 
 * 
 * MyCircularDeque(k)：构造函数,双端队列的大小为k。
 * insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true。
 * insertLast()：将一个元素添加到双端队列尾部。如果操作成功返回 true。
 * deleteFront()：从双端队列头部删除一个元素。 如果操作成功返回 true。
 * deleteLast()：从双端队列尾部删除一个元素。如果操作成功返回 true。
 * getFront()：从双端队列头部获得一个元素。如果双端队列为空，返回 -1。
 * getRear()：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1。
 * isEmpty()：检查双端队列是否为空。
 * isFull()：检查双端队列是否满了。
 * 
 * 
 * 示例：
 * 
 * MyCircularDeque circularDeque = new MycircularDeque(3); // 设置容量大小为3
 * circularDeque.insertLast(1);                    // 返回 true
 * circularDeque.insertLast(2);                    // 返回 true
 * circularDeque.insertFront(3);                    // 返回 true
 * circularDeque.insertFront(4);                    // 已经满了，返回 false
 * circularDeque.getRear();                  // 返回 2
 * circularDeque.isFull();                        // 返回 true
 * circularDeque.deleteLast();                    // 返回 true
 * circularDeque.insertFront(4);                    // 返回 true
 * circularDeque.getFront();                // 返回 4
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 所有值的范围为 [1, 1000]
 * 操作次数的范围为 [1, 1000]
 * 请不要使用内置的双端队列库。
 * 
 * 
 */
public class MyCircularDeque {

    class Node 
    {
        public int Value;
        public Node Next;

        public Node Prev;

        public Node(int val=0, Node next=null, Node prev=null)
        {
            this.Value = val;
            this.Next = next;
            this.Prev = prev;
        }

        public void InserAfter(Node node)
        {
            node.Next = this.Next;
            this.Next = node;
            node.Prev = this;
            node.Next.Prev = node;
        }

        public void InserBefore(Node node)
        {
            node.Prev = this.Prev;
            this.Prev = node;
            node.Prev.Next = node;
            node.Next = this;
        }

        public static void Remove(Node node)
        {
            node.Prev.Next = node.Next;
            node.Next.Prev = node.Prev;
        }
    }

    private Node header;
    private Node tailer;

    private int capacity;

    private int count;

    /** Initialize your data structure here. Set the size of the deque to be k. */
    public MyCircularDeque(int k) {
        this.capacity = k;
        this.count = 0;
        this.header = new Node();
        this.tailer = new Node();
        this.header.Next = this.tailer;
        this.tailer.Prev = this.header;
    }
    
    /** Adds an item at the front of Deque. Return true if the operation is successful. */
    public bool InsertFront(int value) {
        if(this.count >= this.capacity)
        {
            return false;
        }
        Node node = new Node(value);
        this.header.InserAfter(node);
        this.count ++;
        return true;
    }
    
    /** Adds an item at the rear of Deque. Return true if the operation is successful. */
    public bool InsertLast(int value) {
        if(this.count >= this.capacity)
        {
            return false;
        }
        Node node = new Node(value);
        this.tailer.InserBefore(node);
        this.count++;
        return true;
    }
    
    /** Deletes an item from the front of Deque. Return true if the operation is successful. */
    public bool DeleteFront() {
        if(this.count == 0)
        {
            return false;
        }
        Node node = this.header.Next;
        Node.Remove(node);
        this.count--;
        return true;
    }
    
    /** Deletes an item from the rear of Deque. Return true if the operation is successful. */
    public bool DeleteLast() {
        if(this.count == 0)
        {
            return false;
        }
        Node node = this.tailer.Prev;
        Node.Remove(node);
        this.count -- ;
        return true;
    }
    
    /** Get the front item from the deque. */
    public int GetFront() {
        if(this.count == 0)
        {
            return -1;
        }
        return this.header.Next.Value;
    }
    
    /** Get the last item from the deque. */
    public int GetRear() {
        if(this.count == 0)
        {
            return -1;
        }
        return this.tailer.Prev.Value;
    }
    
    /** Checks whether the circular deque is empty or not. */
    public bool IsEmpty() {
        return this.count == 0 ;
    }
    
    /** Checks whether the circular deque is full or not. */
    public bool IsFull() {
        return this.count == this.capacity;
    }
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * MyCircularDeque obj = new MyCircularDeque(k);
 * bool param_1 = obj.InsertFront(value);
 * bool param_2 = obj.InsertLast(value);
 * bool param_3 = obj.DeleteFront();
 * bool param_4 = obj.DeleteLast();
 * int param_5 = obj.GetFront();
 * int param_6 = obj.GetRear();
 * bool param_7 = obj.IsEmpty();
 * bool param_8 = obj.IsFull();
 */

