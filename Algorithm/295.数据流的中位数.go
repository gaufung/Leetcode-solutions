/*
 * @lc app=leetcode.cn id=295 lang=golang
 *
 * [295] 数据流的中位数
 *
 * https://leetcode-cn.com/problems/find-median-from-data-stream/description/
 *
 * algorithms
 * Hard (34.54%)
 * Likes:    44
 * Dislikes: 0
 * Total Accepted:    2.8K
 * Total Submissions: 8.1K
 * Testcase Example:  '["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]\n[[],[1],[2],[],[3],[]]'
 *
 * 中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。
 *
 * 例如，
 *
 * [2,3,4] 的中位数是 3
 *
 * [2,3] 的中位数是 (2 + 3) / 2 = 2.5
 *
 * 设计一个支持以下两种操作的数据结构：
 *
 *
 * void addNum(int num) - 从数据流中添加一个整数到数据结构中。
 * double findMedian() - 返回目前所有元素的中位数。
 *
 *
 * 示例：
 *
 * addNum(1)
 * addNum(2)
 * findMedian() -> 1.5
 * addNum(3)
 * findMedian() -> 2
 *
 * 进阶:
 *
 *
 * 如果数据流中所有整数都在 0 到 100 范围内，你将如何优化你的算法？
 * 如果数据流中 99% 的整数都在 0 到 100 范围内，你将如何优化你的算法？
 *
 *
 */
type Heap struct {
	elements []int
	compare  func(int, int) bool
}

func NewHeap(compare func(int, int) bool) *Heap {
	return &Heap{
		elements: make([]int, 0),
		compare:  compare,
	}
}

func (h *Heap) heap() {
	i := len(h.elements) - 1
	for ; i > 0; i-- {
		j := (i - 1) / 2
		if j >= 0 && h.compare(h.elements[i], h.elements[j]) {
			h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
		}
	}

}

func (h *Heap) upHeap() {
	i := len(h.elements) - 1
	for i > 0 {
		j := (i - 1) / 2
		if h.compare(h.elements[i], h.elements[j]) {
			h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
			i = j
		} else {
			break
		}
	}
}

func (h *Heap) downHeap() {
	size := len(h.elements)
	i := 0
	for i < size {
		left, right := 2*i+1, 2*i+2
		if right < size {
			index := right
			if h.compare(h.elements[left], h.elements[right]) {
				index = left
			}
			if h.compare(h.elements[index], h.elements[i]) {
				h.elements[index], h.elements[i] = h.elements[i], h.elements[index]
				i = index
			} else {
				break
			}
		} else if left < size {
			if h.compare(h.elements[left], h.elements[i]) {
				h.elements[left], h.elements[i] = h.elements[i], h.elements[left]
				i = left
				continue
			} else {
				break
			}
		} else {
			break
		}
	}
}

func (h *Heap) Top() int {
	return h.elements[0]
}
func (h *Heap) Pop() int {
	val := h.Top()
	h.elements[0], h.elements[len(h.elements)-1] = h.elements[len(h.elements)-1], h.elements[0]
	h.elements = h.elements[:len(h.elements)-1]
	h.downHeap()
	return val
}

func (h *Heap) Push(val int) {
	h.elements = append(h.elements, val)
	h.upHeap()
}

func (h *Heap) Empty() bool {
	return len(h.elements) == 0
}

type MedianFinder struct {
	minHeap *Heap
	maxHeap *Heap
	cnt     int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	min := func(a, b int) bool { return a < b }
	max := func(a, b int) bool { return a > b }
	finder := MedianFinder{cnt: 0}
	finder.minHeap = NewHeap(min)
	finder.maxHeap = NewHeap(max)
	return finder
}

func (this *MedianFinder) AddNum(num int) {
	if this.cnt%2 == 0 {
		this.maxHeap.Push(num)
	} else {
		this.minHeap.Push(num)
	}
	this.cnt = this.cnt + 1
	if this.minHeap.Empty() || this.maxHeap.Empty() {
		return
	}
	min := this.minHeap.Top()
	max := this.maxHeap.Top()
	if max > min {
		this.minHeap.Pop()
		this.minHeap.Push(max)
		this.maxHeap.Pop()
		this.maxHeap.Push(min)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.cnt%2 == 0 {
		return (float64(this.minHeap.Top()) + float64(this.maxHeap.Top())) * 0.5
	} else {
		return float64(this.maxHeap.Top())
	}
}

