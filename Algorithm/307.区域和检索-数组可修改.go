/*
 * @lc app=leetcode.cn id=307 lang=golang
 *
 * [307] 区域和检索 - 数组可修改
 *
 * https://leetcode-cn.com/problems/range-sum-query-mutable/description/
 *
 * algorithms
 * Medium (49.65%)
 * Likes:    48
 * Dislikes: 0
 * Total Accepted:    2.5K
 * Total Submissions: 5.1K
 * Testcase Example:  '["NumArray","sumRange","update","sumRange"]\n[[[1,3,5]],[0,2],[1,2],[0,2]]'
 *
 * 给定一个整数数组  nums，求出数组从索引 i 到 j  (i ≤ j) 范围内元素的总和，包含 i,  j 两点。
 *
 * update(i, val) 函数可以通过将下标为 i 的数值更新为 val，从而对数列进行修改。
 *
 * 示例:
 *
 * Given nums = [1, 3, 5]
 *
 * sumRange(0, 2) -> 9
 * update(1, 2)
 * sumRange(0, 2) -> 8
 *
 *
 * 说明:
 *
 *
 * 数组仅可以在 update 函数下进行修改。
 * 你可以假设 update 函数与 sumRange 函数的调用次数是均匀分布的。
 *
 *
 */
type NumArray struct {
	subSum []int
	nums   []int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	if n > 0 {
		subSum := make([]int, n)
		subSum[0] = nums[0]
		for i := 1; i < n; i++ {
			subSum[i] = subSum[i-1] + nums[i]
		}
		return NumArray{
			subSum: subSum,
			nums:   nums,
		}
	} else {
		return NumArray{}
	}

}

func (this *NumArray) Update(i int, val int) {
	diff := val - this.nums[i]
	this.nums[i] = val
	for i < len(this.nums) {
		this.subSum[i] += diff
		i++
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.subSum[j]
	} else {
		return this.subSum[j] - this.subSum[i-1]
	}
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */

