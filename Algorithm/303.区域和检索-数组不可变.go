/*
 * @lc app=leetcode.cn id=303 lang=golang
 *
 * [303] 区域和检索 - 数组不可变
 *
 * https://leetcode-cn.com/problems/range-sum-query-immutable/description/
 *
 * algorithms
 * Easy (53.32%)
 * Likes:    73
 * Dislikes: 0
 * Total Accepted:    11.1K
 * Total Submissions: 20.9K
 * Testcase Example:  '["NumArray","sumRange","sumRange","sumRange"]\n[[[-2,0,3,-5,2,-1]],[0,2],[2,5],[0,5]]'
 *
 * 给定一个整数数组  nums，求出数组从索引 i 到 j  (i ≤ j) 范围内元素的总和，包含 i,  j 两点。
 *
 * 示例：
 *
 * 给定 nums = [-2, 0, 3, -5, 2, -1]，求和函数为 sumRange()
 *
 * sumRange(0, 2) -> 1
 * sumRange(2, 5) -> -1
 * sumRange(0, 5) -> -3
 *
 * 说明:
 *
 *
 * 你可以假设数组不可变。
 * 会多次调用 sumRange 方法。
 *
 *
 */
type NumArray struct {
	partialSum []int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	if n <= 0 {
		return NumArray{
			partialSum: make([]int, 0),
		}
	}
	partialSum := make([]int, n)
	partialSum[0] = nums[0]
	for i := 1; i < n; i++ {
		partialSum[i] += partialSum[i-1] + nums[i]
	}
	return NumArray{
		partialSum: partialSum,
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	n := len(this.partialSum)
	if j >= n {
		j = n - 1
	}
	if i > 0 {
		return this.partialSum[j] - this.partialSum[i-1]
	} else {
		return this.partialSum[j]
	}
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

