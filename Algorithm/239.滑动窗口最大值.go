/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 *
 * https://leetcode-cn.com/problems/sliding-window-maximum/description/
 *
 * algorithms
 * Hard (40.81%)
 * Likes:    83
 * Dislikes: 0
 * Total Accepted:    6.7K
 * Total Submissions: 16.4K
 * Testcase Example:  '[1,3,-1,-3,5,3,6,7]\n3'
 *
 * 给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口 k
 * 内的数字。滑动窗口每次只向右移动一位。
 *
 * 返回滑动窗口最大值。
 *
 * 示例:
 *
 * 输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
 * 输出: [3,3,5,5,6,7]
 * 解释:
 *
 * ⁠ 滑动窗口的位置                最大值
 * ---------------               -----
 * [1  3  -1] -3  5  3  6  7       3
 * ⁠1 [3  -1  -3] 5  3  6  7       3
 * ⁠1  3 [-1  -3  5] 3  6  7       5
 * ⁠1  3  -1 [-3  5  3] 6  7       5
 * ⁠1  3  -1  -3 [5  3  6] 7       6
 * ⁠1  3  -1  -3  5 [3  6  7]      7
 *
 * 注意：
 *
 * 你可以假设 k 总是有效的，1 ≤ k ≤ 输入数组的大小，且输入数组不为空。
 *
 * 进阶：
 *
 * 你能在线性时间复杂度内解决此题吗？
 *
 */
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n <= 0 {
		return []int{}
	}
	result := make([]int, n-k+1)
	q := NewQueue()
	for i := 0; i < n; i++ {
		for q.Size() > 0 && i-k == q.Front() {
			q.DequeueFront()
		}
		for q.Size() > 0 && nums[q.Back()] < nums[i] {
			q.DequeueBack()
		}
		q.EnqueueBack(i)
		if i+1-k >= 0 {
			result[i+1-k] = nums[q.Front()]
		}
	}
	return result
}

type Queue struct {
	elements []int
}

func NewQueue() *Queue {
	return &Queue{
		elements: make([]int, 0),
	}
}
func (q *Queue) Size() int {
	return len(q.elements)
}
func (q *Queue) DequeueFront() int {
	val := q.elements[0]
	q.elements = q.elements[1:]
	return val
}
func (q *Queue) DequeueBack() int {
	val := q.elements[q.Size()-1]
	q.elements = q.elements[:q.Size()-1]
	return val
}

func (q *Queue) EnqueueBack(val int) {
	q.elements = append(q.elements, val)
}

func (q *Queue) Front() int {
	return q.elements[0]
}
func (q *Queue) Back() int {
	return q.elements[len(q.elements)-1]
}


