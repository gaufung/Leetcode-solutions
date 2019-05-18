import "sort"

/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 *
 * https://leetcode-cn.com/problems/4sum/description/
 *
 * algorithms
 * Medium (35.22%)
 * Likes:    224
 * Dislikes: 0
 * Total Accepted:    18.8K
 * Total Submissions: 53.4K
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * 给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c
 * + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
 *
 * 注意：
 *
 * 答案中不可以包含重复的四元组。
 *
 * 示例：
 *
 * 给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
 *
 * 满足要求的四元组集合为：
 * [
 * ⁠ [-1,  0, 0, 1],
 * ⁠ [-2, -1, 1, 2],
 * ⁠ [-2,  0, 0, 2]
 * ]
 *
 *
 */
func fourSum(nums []int, target int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		} else {
			a := nums[i]
			for j := i + 1; j < len(nums)-2; j++ {
				if j-1 > i && nums[j-1] == nums[j] {
					continue
				} else {
					b := nums[j]
					l, k := j+1, len(nums)-1
					for l < k {
						sum := a + b + nums[l] + nums[k]
						if sum > target {
							k--
						} else if sum < target {
							l++
						} else {
							item := []int{a, b, nums[l], nums[k]}
							result = append(result, item)
							l++
							k--
							for nums[k] == nums[k+1] && nums[l] == nums[l-1] && l < k {
								l++
								k--
							}
						}
					}
				}
			}
		}
	}
	return result
}

