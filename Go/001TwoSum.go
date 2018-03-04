func twoSum(nums []int, target int) []int {
    var dic = make(map[int] []int)
    for idx, num := range(nums){
        _, ok := dic[num]
        if !ok {
            dic[num] = make([]int,0)
        }
        dic[num] = append(dic[num], idx)
    }
    for _, num := range(nums){
        left := target - num
        if _, ok := dic[left]; ok {
            if left != num {
                return []int { dic[num][0], dic[left][0] }
            }else{
                if len(dic[num]) > 1 {
                    return []int{dic[num][0], dic[num][1] }
                }
            }
        }
    }
    return nil
}