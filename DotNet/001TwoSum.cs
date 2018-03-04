public class Solution {
    public int[] TwoSum(int[] nums, int target) {
        IDictionary<int, IList<int>> dic = new Dictionary<int, IList<int>>();
        for(int i=0; i<nums.Length;i++){
            if(!dic.ContainsKey(nums[i]))
                dic[nums[i]] = new List<int>();
            dic[nums[i]].Add(i);
        }
        foreach(int num in nums){
            int left = target - num;
            if(dic.ContainsKey(left)){
                if(left!=num)
                    return new []{dic[num][0], dic[left][0]};
                else{
                    if(dic[num].Count>1)
                        return new []{dic[num][0], dic[left][1]};
                }
            }
        }
        return new int[]{};  
    }
}