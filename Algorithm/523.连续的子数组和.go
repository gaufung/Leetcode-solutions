/*
 * @lc app=leetcode.cn id=523 lang=golang
 *
 * [523] 连续的子数组和
 */
func checkSubarraySum(nums []int, k int) bool {
	n := len(nums)
	partialSum := make([]int, n)
	partialSum[0] = nums[0]
	for i := 1; i < n; i++ {
		partialSum[i] = partialSum[i-1] + nums[i]
	}
	for i := 1; i < n; i++ {
		if k == 0 {
			if partialSum[i] == 0 {
				return true
			}
			for j := i - 2; j >= 0; j-- {
				if partialSum[i]-partialSum[j] == 0 {
					return true
				}
			}
		} else {
			if partialSum[i]%k == 0 {
				return true
			}
			for j := i - 2; j >= 0; j-- {
				if (partialSum[i]-partialSum[j])%k == 0 {
					return true
				}
			}
		}

	}
	return false

}

