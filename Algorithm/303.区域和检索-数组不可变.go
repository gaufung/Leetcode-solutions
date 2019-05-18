/*
 * @lc app=leetcode.cn id=303 lang=golang
 *
 * [303] 区域和检索 - 数组不可变
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

