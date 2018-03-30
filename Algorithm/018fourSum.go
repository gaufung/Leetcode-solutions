func fourSum(nums []int, target int) [][]int {
    result := make([][]int, 0)
    sort.Ints(nums)
    for i:=0; i<len(nums)-3;i++{
        if i>0 && nums[i-1] == nums[i] {
            continue
        }else{
            a := nums[i]
            for j:=i+1; j < len(nums)-2;j++ {
                if j-1 > i && nums[j-1] == nums[j] {
                    continue
                }else{
                    b := nums[j]
                    l,k := j+1, len(nums)-1
                    for l < k {
                        sum:= a + b + nums[l] + nums[k]
                        if sum > target {
                            k--
                        }else if sum < target {
                            l++
                        }else{
                            item := []int{a,b,nums[l],nums[k]}
                            result = append(result, item)
                            l++
                            k--
                            for nums[k]==nums[k+1] && nums[l]==nums[l-1] && l<k{
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