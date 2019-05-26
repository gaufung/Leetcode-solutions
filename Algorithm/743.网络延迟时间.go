/*
 * @lc app=leetcode.cn id=743 lang=golang
 *
 * [743] 网络延迟时间
 *
 * https://leetcode-cn.com/problems/network-delay-time/description/
 *
 * algorithms
 * Medium (42.87%)
 * Likes:    33
 * Dislikes: 0
 * Total Accepted:    1.7K
 * Total Submissions: 4K
 * Testcase Example:  '[[2,1,1],[2,3,1],[3,4,1]]\n4\n2'
 *
 * 有 N 个网络节点，标记为 1 到 N。
 *
 * 给定一个列表 times，表示信号经过有向边的传递时间。 times[i] = (u, v, w)，其中 u 是源节点，v 是目标节点， w
 * 是一个信号从源节点传递到目标节点的时间。
 *
 * 现在，我们向当前的节点 K 发送了一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回 -1。
 *
 * 注意:
 *
 *
 * N 的范围在 [1, 100] 之间。
 * K 的范围在 [1, N] 之间。
 * times 的长度在 [1, 6000] 之间。
 * 所有的边 times[i] = (u, v, w) 都有 1 <= u, v <= N 且 0 <= w <= 100。
 *
 *
 */
func networkDelayTime(times [][]int, N int, K int) int {
	visited := make([]bool, N+1)
	timestamps := make([]int, N+1)
	h := NewHeap()
	h.Add(NewNode(K, 0))
	for h.Size() > 0 {
		start := h.Pop()
		visited[start.Label] = true
		timestamp := start.Time
		timestamps[start.Label] = timestamp
		for _, uvw := range times {
			if uvw[0] == start.Label && visited[uvw[1]] == false {
				if node := h.Lookup(uvw[1]); node != nil {
					node.Time = min(node.Time, timestamp+uvw[2])
				} else {
					node = NewNode(uvw[1], timestamp+uvw[2])
					h.Add(node)
				}
			}
		}
		h.heap()
	}
	val := 0
	for i := 1; i < N+1; i++ {
		if visited[i] == false {
			val = -1
			break
		} else {
			val = max(val, timestamps[i])
		}
	}
	return val
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Node struct {
	Label int
	Time  int
}

func NewNode(label, time int) *Node {
	return &Node{
		Label: label,
		Time:  time,
	}
}

type Heap struct {
	Nodes  []*Node
	LookUp map[int]*Node
}

func NewHeap() *Heap {
	return &Heap{
		Nodes:  make([]*Node, 0),
		LookUp: make(map[int]*Node, 0),
	}
}

func (h *Heap) heap() {
	for i := len(h.Nodes) - 1; i > 0; i-- {
		if h.Nodes[i].Time < h.Nodes[(i-1)/2].Time {
			h.Nodes[i], h.Nodes[(i-1)/2] = h.Nodes[(i-1)/2], h.Nodes[i]
		}
	}
}

func (h *Heap) Add(node *Node) {
	h.Nodes = append(h.Nodes, node)
	h.LookUp[node.Label] = node
}

func (h *Heap) Lookup(label int) *Node {
	return h.LookUp[label]
}

func (h *Heap) Pop() *Node {
	n := h.Size()
	val := h.Nodes[0]
	h.Nodes[0], h.Nodes[n-1] = h.Nodes[n-1], h.Nodes[0]
	h.Nodes = h.Nodes[:n-1]
	delete(h.LookUp, val.Label)
	return val
}

func (h *Heap) heapDown() {
	n := h.Size()
	if n <= 0 {
		return
	}
	i := 0
	for i < n {
		left := 2*i + 1
		right := 2*i + 2
		if right < n {
			index := left
			if h.Nodes[right].Time < h.Nodes[left].Time {
				index = right
			}
			if h.Nodes[index].Time < h.Nodes[i].Time {
				h.Nodes[index], h.Nodes[i] = h.Nodes[i], h.Nodes[index]
				i = index
			} else {
				break
			}
		} else if left < n {
			if h.Nodes[left].Time < h.Nodes[i].Time {
				h.Nodes[left], h.Nodes[i] = h.Nodes[i], h.Nodes[left]
				i = left
			} else {
				break
			}
		} else {
			break
		}
	}
}

func (h *Heap) Size() int {
	return len(h.Nodes)
}

