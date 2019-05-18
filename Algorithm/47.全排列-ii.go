import "sort"

/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */
func permuteUnique(nums []int) [][]int {
	n := len(nums)
	if n <= 0 {
		return [][]int{}
	}
	if n == 1 {
		return [][]int{
			[]int{nums[0]},
		}
	}
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		if i != len(nums)-1 && nums[i] == nums[i+1] {
			continue
		}
		leftResult := permuteUnique(removeKth(nums, i))
		for _, items := range leftResult {
			result = append(result, append(items, nums[i]))
		}
	}
	return result
}
func removeKth(nums []int, k int) []int {
	result := make([]int, 0)
	for i, val := range nums {
		if i != k {
			result = append(result, val)
		}
	}
	return result
}

