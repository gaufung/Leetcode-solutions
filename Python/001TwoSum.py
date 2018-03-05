class Solution:
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        dic = {}
        for idx, num in enumerate(nums):
            if num not in dic:
                dic[num]=[]
            dic[num].append(idx)
        for num in nums:
            left = target - num
            if left in dic:
                if left != num:
                    return [dic[num][0], dic[left][0]]
                else:
                    if len(dic[num]) > 1:
                        return [dic[num][0], dic[num][1]]