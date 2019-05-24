import "math/rand"

/*
 * @lc app=leetcode.cn id=384 lang=golang
 *
 * [384] 打乱数组
 *
 * https://leetcode-cn.com/problems/shuffle-an-array/description/
 *
 * algorithms
 * Medium (45.19%)
 * Likes:    25
 * Dislikes: 0
 * Total Accepted:    7.2K
 * Total Submissions: 15.8K
 * Testcase Example:  '["Solution","shuffle","reset","shuffle"]\n[[[1,2,3]],[],[],[]]'
 *
 * 打乱一个没有重复元素的数组。
 *
 * 示例:
 *
 *
 * // 以数字集合 1, 2 和 3 初始化数组。
 * int[] nums = {1,2,3};
 * Solution solution = new Solution(nums);
 *
 * // 打乱数组 [1,2,3] 并返回结果。任何 [1,2,3]的排列返回的概率应该相同。
 * solution.shuffle();
 *
 * // 重设数组到它的初始状态[1,2,3]。
 * solution.reset();
 *
 * // 随机返回数组[1,2,3]打乱后的结果。
 * solution.shuffle();
 *
 *
 */
type Solution struct {
	origin  []int
	shuffle []int
}

func Constructor(nums []int) Solution {
	n := len(nums)
	origin, shuffle := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		origin[i] = nums[i]
		shuffle[i] = nums[i]
	}
	return Solution{
		origin:  origin,
		shuffle: shuffle,
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	copy(this.shuffle, this.origin)
	return this.shuffle
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	for i := 0; i < len(this.origin); i++ {
		swapIndex := rand.Intn(len(this.origin))
		this.shuffle[i], this.shuffle[swapIndex] = this.shuffle[swapIndex], this.shuffle[i]
	}
	return this.shuffle
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

