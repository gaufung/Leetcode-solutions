import "sort"

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 *
 * https://leetcode-cn.com/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (40.48%)
 * Likes:    184
 * Dislikes: 0
 * Total Accepted:    25.4K
 * Total Submissions: 62.8K
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * 给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target
 * 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
 *
 * 例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.
 *
 * 与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
 *
 *
 */
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	if len(nums) <= 3 {
		result := 0
		for _, val := range nums {
			result += val
		}
		return result
	}
	minTarget := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			currentTarget := nums[i] + nums[j] + nums[k]
			if currentTarget == target {
				return target
			}
			if currentTarget > target {
				if abs(currentTarget-target) < abs(minTarget-target) {
					minTarget = currentTarget
				}
				k--
				for nums[k] == nums[k+1] && j < k {
					k--
				}
			} else {
				if abs(currentTarget-target) < abs(minTarget-target) {
					minTarget = currentTarget
				}
				j++
				for nums[j] == nums[j-1] && j < k {
					j++
				}
			}
		}
	}
	return minTarget
}

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

