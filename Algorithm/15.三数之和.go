import "sort"

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 *
 * https://leetcode-cn.com/problems/3sum/description/
 *
 * algorithms
 * Medium (22.34%)
 * Likes:    982
 * Dislikes: 0
 * Total Accepted:    55.3K
 * Total Submissions: 247.5K
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0
 * ？找出所有满足条件且不重复的三元组。
 *
 * 注意：答案中不可以包含重复的三元组。
 *
 * 例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
 *
 * 满足要求的三元组集合为：
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 *
 *
 */
func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	if nums == nil || len(nums) < 3 {
		return result
	}
	sort.Ints(nums)
	if nums[0] > 0 || nums[len(nums)-1] < 0 {
		return result
	}
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else {
				if nums[i]+nums[j]+nums[k] < 0 {
					j++
				} else {
					item := []int{nums[i], nums[j], nums[k]}
					result = append(result, item)
					j++
					k--
					for nums[j] == nums[j-1] && nums[k] == nums[k+1] && j < k {
						j++
						k--
					}
				}
			}
		}
	}
	return result
}

