func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	if nums == nil || len(nums) < 3 {
		return result
	}
	sort.Ints(nums)
	if nums[0] > 0 || nums[len(nums)-1] < 0 {
		return result
	}
	for i:= 0; i < len(nums)-2; i++ {
		if i>0 && nums[i] == nums[i-1]{
			continue
		}
		j, k := i+1, len(nums) - 1
		for j < k {
			if nums[i] + nums[j] + nums[k] > 0 {
				k--
			}else{
				if nums[i] + nums[j] + nums[k] < 0 {
					j++
				}else{
					item := []int{nums[i], nums[j], nums[k]}
					result = append(result, item)
					j++
					k--
					for nums[j]==nums[j-1] && nums[k] == nums[k+1] && j < k {
						j++
						k--
					}
				}
			}
		}
	}
	return result
}